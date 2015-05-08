package main

import (
	"strings"
)

func GitConfig(args ...string) (string, error) {
	gitArgs := append([]string{"config", "--path", "--null"}, args...)
	buf, err := RunWithOutput("git", gitArgs...)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(buf), "\000"), nil
}

func GitConfigSingle(key string) (string, error) {
	return GitConfig("--get", key)
}
