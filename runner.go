package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	// return cmd.Run()
}
