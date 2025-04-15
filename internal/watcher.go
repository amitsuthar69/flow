package internal

import (
	"io/fs"
	"os"
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
		Log("error", err.Error())
	}
	defer watcher.Close()

	if err := Walk(watcher, cfg.Root, cfg.Build.ExcludeDir, cfg.Build.ExcludeRegex); err != nil {
		Log("fatal", err.Error())
	}

	err = os.MkdirAll("tmp", 0755)
	if err != nil {
		Log("fatal", err.Error())
	}

	var (
		timer *time.Timer
	)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {

					// regex pattern check
					fileName := filepath.Base(event.Name)
					excluded := false
					for _, pattern := range cfg.Build.ExcludeRegex {
						if ok, _ := filepath.Match(pattern, fileName); ok {
							excluded = true
							break
						}
					}
					if excluded {
						continue
					}

					ext := strings.TrimPrefix(filepath.Ext(event.Name), ".")
					if !slices.Contains(cfg.Build.IncludeExt, ext) {
						continue
					}

					if timer != nil {
						timer.Stop()
					}

					var debounce int
					if cfg.Debounce > 0 {
						debounce = cfg.Debounce
					} else {
						debounce = 500
					}

					timer = time.AfterFunc(time.Duration(debounce)*time.Millisecond, func() {
						Build(cfg.Build.Cmd, event.Name)
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				Log("error", err.Error())
			}
		}
	}()

	// block forever...
	select {}
}

func Walk(watcher *fsnotify.Watcher, root string, excludeDirs []string, excludeRegex []string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// exclude directories from excludeDirs
			for _, exclude := range excludeDirs {
				if path == ".git" || path == "tmp" || strings.Contains(path, exclude) {
					return filepath.SkipDir
				}
			}
			if err := watcher.Add(path); err != nil {
				return err
			}
			return nil
		}

		return nil
	})
}
