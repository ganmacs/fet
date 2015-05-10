package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type LocalRepository struct {
	FullPath  string
	RelPath   string
	PathParts []string
}

func (repo *LocalRepository) Match(str string) bool {
	for _, p := range repo.SubPaths() {
		if p == str {
			return true
		}
	}

	return false
}

func (repo *LocalRepository) SubPaths() []string {
	sub := make([]string, len(repo.PathParts))
	for i := range repo.PathParts {
		sub[i] = filepath.Join(repo.PathParts[i:]...)
	}
	return sub
}

var _rootRepositry string

func LocalRepositoryRoot() string {
	if _rootRepositry != "" {
		return _rootRepositry
	}

	_rootRepositry, err := GitConfigSingle("ghq.root")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return _rootRepositry
}

func NewLocalRepositoryFromURL(remote *url.URL) *LocalRepository {
	pathParts := append([]string{remote.Host}, strings.Split(remote.Path, "/")...)
	relPath := strings.TrimSuffix(path.Join(pathParts...), ".git")

	return &LocalRepository{
		path.Join(LocalRepositoryRoot(), relPath),
		relPath,
		pathParts,
	}
}

func NewLocalRepositoryFromFullPath(fullPath string) (*LocalRepository, error) {
	relPath, err := filepath.Rel(LocalRepositoryRoot(), fullPath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &LocalRepository{
		fullPath,
		relPath,
		strings.Split(relPath, "/"),
	}, nil
}

func WalkLocalRepositories(root string, callback func(*LocalRepository)) error {
	return filepath.Walk(root, func(fullPath string, fileInfo os.FileInfo, err error) error {
		if fileInfo == nil || !fileInfo.IsDir() {
			return nil
		}

		_, err = os.Stat(filepath.Join(fullPath, ".git"))
		if err != nil {
			return nil
		}

		// TODO Don't use struct
		local, err := NewLocalRepositoryFromFullPath(fullPath)
		if err != nil {
			return nil
		}

		if local == nil {
			return nil // TODO
		}

		callback(local)

		return filepath.SkipDir
	})
}
