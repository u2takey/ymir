package main

import (
	"fmt"
	"os"

	"github.com/arlert/ymir/cmd/agent"
	"github.com/arlert/ymir/cmd/server"
	"github.com/arlert/ymir/version"

	"github.com/ianschenck/envflag"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

func main() {
	envflag.Parse()

	app := cli.NewApp()
	app.Name = "ymir"
	app.Version = version.Version.String()
	app.Usage = "command line utility"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		server.Command,
		agent.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
