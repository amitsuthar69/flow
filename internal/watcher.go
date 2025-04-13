package internal

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watch() {
	cfg := ParseTomlConfig()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
	}
	defer watcher.Close()

	if err := Walk(watcher, cfg.Root, cfg.Build.ExcludeDir); err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("tmp", 0755)
	if err != nil {
		log.Fatal(err)
	}

	var (
		timer     *time.Timer
		lastEvent string
	)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					ext := strings.TrimPrefix(filepath.Ext(event.Name), ".")
					if !slices.Contains(cfg.Build.IncludeExt, ext) {
						continue
					}

					if timer != nil {
						timer.Stop()
					}

					lastEvent = event.Name
					timer = time.AfterFunc(500*time.Millisecond, func() {
						Build(cfg.Build.Cmd, lastEvent)
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// block forever...
	select {}
}

func Walk(watcher *fsnotify.Watcher, root string, excludeDirs []string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// if not a directory, we skip
		if !info.IsDir() {
			return nil
		}

		// exclude directories from excludeDirs
		for _, exclude := range excludeDirs {
			if path == ".git" {
				return filepath.SkipDir
			}

			if path == "tmp" {
				return filepath.SkipDir
			}

			if strings.Contains(path, exclude) {
				return filepath.SkipDir
			}
		}

		// add path to watcher
		if err := watcher.Add(path); err != nil {
			return err
		} else {
			log.Printf("watching: '%s'\n", path)
		}

		return nil
	})
}

func Build(buildCmd string, file string) {
	log.Printf("file %s has been modified, rebuilding...", file)
	go func() {
		cmd := exec.Command("sh", "-c", buildCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Printf("build error: %v", err)
		}
	}()
}
