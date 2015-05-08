package main

import (
	"os"
	"os/exec"
)

func Run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunWithOutput(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}
