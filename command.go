package main

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
	"os"
	"syscall"
)

var Commands = []cli.Command{
	GetCommand,
	LookCommand,
	ListCommand,
}

var GetCommand = cli.Command{
	Name:        "get",
	Usage:       "clone repository",
	Description: "this is desc",
	Action:      doGetCommand,
}

var LookCommand = cli.Command{
	Name:        "look",
	Usage:       "look into a local repository",
	Description: "this is desc",
	Action:      doLookCommand,
}

var ListCommand = cli.Command{
	Name:        "list",
	Usage:       "list all local repostory",
	Description: "this is desc",
	Action:      doListCommand,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "full-path, p", Usage: "Print full paths"},
	},
}

func doGetCommand(c *cli.Context) {
	argUrl := c.Args().Get(0)
	url, err := NewUrl(argUrl)
	if err != nil {
		panic(err)
	}

	remote, err := NewRemoteRepository(url)
	if err != nil {
		panic(err)
	}

	if !remote.isValid() {
		err = errors.New(fmt.Sprintf("Not a Valid repository: %s", url))
		fmt.Println(err)
	}

	remote.GetRepositroy()
}

func doLookCommand(c *cli.Context) {
	relPath := c.Args().Get(0)

	foundedRepo := []*LocalRepository{}
	err := WalkLocalRepositories(func(repo *LocalRepository) {
		if repo.Match(relPath) {
			foundedRepo = append(foundedRepo, repo)
		}
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	switch len(foundedRepo) {
	case 0:
		fmt.Println("not found")
		return
	case 1:
		shell := os.Getenv("SHELL")

		fmt.Printf("cd %s\n", foundedRepo[0].FullPath)
		err := os.Chdir(foundedRepo[0].FullPath)
		if err != nil {
			fmt.Println(err)
		}
		syscall.Exec(shell, []string{shell}, syscall.Environ())
	default:
		fmt.Println("not invalid")
		return
	}
}

func doListCommand(c *cli.Context) {
	isFullPath := c.Bool("full-path")

	repos := []*LocalRepository{}
	err := WalkLocalRepositories(func(repo *LocalRepository) {
		repos = append(repos, repo)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range repos {
		if isFullPath {
			fmt.Println(r.FullPath)
		} else {
			fmt.Println(r.RelPath)
		}
	}
}
