package main

import (
	"net/url"
	"path"
	"strings"
)

type LocalRepository struct {
	FullPath string
	RelPath  string
}

func NewLocalRepositoryFromURL(remote *url.URL) *LocalRepository {
	pathParts := append([]string{remote.Host}, strings.Split(remote.Path, "/")...)
	relPath := strings.TrimSuffix(path.Join(pathParts...), ".git")

	root, err := GitConfigSingle("ghq.root")
	if err != nil {
		panic(err)
	}

	return &LocalRepository{
		path.Join(root, relPath),
		relPath,
	}
}
