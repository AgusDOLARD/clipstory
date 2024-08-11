package commands

import (
	"fmt"

	"github.com/AgusDOLARD/clipstory/socket"
	"github.com/urfave/cli/v2"
)

func (command *Command) List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "List all clips",
		Action: func(c *cli.Context) error {
			clips, err := socket.GetClips(command.SocketPath)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			for _, clip := range clips {
				fmt.Println(clip)
			}
			return nil
		},
	}
}
