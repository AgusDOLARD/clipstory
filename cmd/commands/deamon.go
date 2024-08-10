package commands

import (
	"github.com/AgusDOLARD/clipstory/socket"
	"github.com/urfave/cli/v2"
)

func (command *Command) Deamon() *cli.Command {
	return &cli.Command{
		Name:    "deamon",
		Aliases: []string{"d"},
		Usage:   "Starts the daemon",
		Action: func(c *cli.Context) error {
			return socket.NewSocket(command.SocketPath).Start()
		},
	}
}
