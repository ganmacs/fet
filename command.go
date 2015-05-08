package main

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	GetCommand,
}

var GetCommand = cli.Command{
	Name:        "get",
	Usage:       "clone repository",
	Description: "this is desc",
	Action:      doGetCommand,
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
