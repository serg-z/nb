package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

type configuration struct {
	NBDir string
}

func (config *configuration) load() error {
	dir := os.Getenv("NB_DIR")

	if dir == "" {
		if runtime.GOOS == "windows" {
			dir = os.Getenv("APPDATA")

			if dir == "" {
				dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "nb")
			}

			dir = filepath.Join(dir, "nb")
		} else {
			dir = filepath.Join(os.Getenv("HOME"), ".nb")
		}
	}

	config.NBDir = dir

	return nil
}

func run() error {
	var config configuration

	err := config.load()
	if err != nil {
		return err
	}

	binary, lookErr := exec.LookPath("nb")
	if lookErr != nil {
		return lookErr
	}

	args := os.Args
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		return execErr
	}

	return nil
}

func main() {
	os.Exit(presentError(run()))
}

func presentError(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	return 0
}