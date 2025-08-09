package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenDir(dir string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", dir)
	case "darwin":
		cmd = exec.Command("open", dir)
	case "linux":
		cmd = exec.Command("xdg-open", dir)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}
