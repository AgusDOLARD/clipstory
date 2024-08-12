package commands

import (
	"github.com/AgusDOLARD/clipstory/socket"
	"github.com/AgusDOLARD/clipstory/tui"
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

			var clip string
			err = tui.Pick(clips, &clip)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return copyToClipboard([]byte(clip))
		},
	}
}

func copyToClipboard(b []byte) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	<-clipboard.Write(clipboard.FmtText, b)
	return nil
}
