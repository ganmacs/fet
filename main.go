package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app().Run(os.Args)
}

func app() *cli.App {
	app := cli.NewApp()
	app.Name = "fetecher"
	app.Usage = "fight the loneliness"
	app.Commands = Commands
	return app
}
