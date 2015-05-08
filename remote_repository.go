package main

import (
	"net/url"
	"strings"
)

type RemoteRepository interface {
	URL() *url.URL
	isValid() bool
	GetRepositroy()
	VCS() VCSBackend
}

type GitHubRepository struct {
	url *url.URL
}

func (remote *GitHubRepository) URL() *url.URL {
	return remote.url
}

func (remote *GitHubRepository) VCS() VCSBackend {
	return GitBackend
}

func (remote *GitHubRepository) isValid() bool {
	pathComponents := strings.Split(strings.TrimRight(remote.url.Path, "/"), "/")
	if len(pathComponents) != 3 {
		return false
	}
	return true
}

func (remote *GitHubRepository) GetRepositroy() {
	remoteURL := remote.URL()
	vcs := remote.VCS()
	vcs.Clone(remoteURL, ".")
	// local := NewLocalRepositoryFromURL(remote)
}

func NewRemoteRepository(url *url.URL) (RemoteRepository, error) {
	if url.Host == "github.com" {
		return &GitHubRepository{url}, nil
	}
	return &GitHubRepository{url}, nil // fixme
}
