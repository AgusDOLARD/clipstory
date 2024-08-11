package commands

import (
	"errors"

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

			if len(clips) == 0 {
				return cli.Exit("No clipboard history", 0)
			}

			var pick string
			err = picker(clips, &pick)
			if err != nil {
				if errors.Is(err, huh.ErrUserAborted) {
					return nil
				}

				return cli.Exit(err.Error(), 1)
			}

			return copyToClipboard([]byte(pick))
		},
	}
}

func picker(clips []string, value *string) error {
	return huh.NewSelect[string]().
		Title("Pick a clip").
		Options(huh.NewOptions(clips...)...).
		// Filtering(true).
		Value(value).
		WithTheme(huh.ThemeBase16()).
		Run()
}

func copyToClipboard(b []byte) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	<-clipboard.Write(clipboard.FmtText, b)
	return nil
}
