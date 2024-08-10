package main

import (
	"fmt"
	"os"

	"github.com/AgusDOLARD/clipstory/cmd/commands"
	"github.com/urfave/cli/v2"
)

func main() {

	comms := commands.Command{SocketPath: "/tmp/clipstory.sock"}
	app := &cli.App{
		Name:  "clipstory",
		Usage: "A simple command line tool to manage clipboard history",
		Commands: []*cli.Command{
			comms.Deamon(),
			comms.List(),
			comms.Pick(),
		},
		Action: func(ctx *cli.Context) error {
			return comms.Pick().Run(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Errorf("Error: %v", err)
	}
}
