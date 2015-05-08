package main

import (
	"net/url"
)

type LocalRepository struct {
	FullPath string
	RelPath  string
}

func NewLocalRepositoryFromURL(remote *url.URL) *LocalRepository {
	return nil
}

func localRepositoryRoots(args) {

}
