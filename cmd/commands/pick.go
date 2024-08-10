package commands

import (
	"github.com/AgusDOLARD/clipstory/socket"
	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func (command *Command) Pick() *cli.Command {
	return &cli.Command{
		Name:    "pick",
		Aliases: []string{"p"},
		Usage:   "Pick a clip",
		Action: func(c *cli.Context) error {
			clips, err := socket.GetClips(command.SocketPath)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			var pick string
			err = huh.NewSelect[string]().
				Title("Pick a clip").
				Options(huh.NewOptions(clips...)...).
				Filtering(true).
				Value(&pick).Run()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			clipboard.Write(clipboard.FmtText, []byte(pick))
			return nil
		},
	}
}
