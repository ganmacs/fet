package main

import (
	"net/url"
)

type VCSBackend interface {
	Clone(*url.URL, string)
}

var GitBackend = &gitBackend{}

type gitBackend struct{}

func (backend *gitBackend) Clone(remote *url.URL, local string) {
	args := []string{"clone"}
	args = append(args, remote.String(), local)
	Run("git", args...)
}
