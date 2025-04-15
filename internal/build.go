package internal

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var (
	currProc *exec.Cmd
	procLock sync.Mutex
)

func Build(buildCmd string, file string) {
	Log("info", fmt.Sprintf("file %s has been modified, rebuilding...", file))

	go func() {

		procLock.Lock()
		// before triggering a new build, we're going to kill the exisiting process
		if currProc != nil && currProc.Process != nil {
			currProc.Process.Kill()
			currProc = nil
		}
		procLock.Unlock()

		cmd := exec.Command("sh", "-c", buildCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			Log("fatal", err.Error())
		}

		procLock.Lock()
		currProc = cmd
		procLock.Unlock()

		err := cmd.Wait()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != -1 {
				// will only log unintended exit errors
				Log("error", err.Error())
			}
		}

		procLock.Lock()
		if currProc == cmd {
			// clear the global process handle after final exit.
			currProc = nil
		}
		procLock.Unlock()

	}()
}
