package commands

import (
	"errors"
	"strings"

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
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "trunc",
				Aliases: []string{"t"},
				Usage:   "Truncate the clip to this length",
				Value:   70,
			},
		},
		Action: func(c *cli.Context) error {
			clips, err := socket.GetClips(command.SocketPath)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			if len(clips) == 0 {
				return cli.Exit("No clipboard history", 0)
			}

			var pick string
			err = picker(clips, &pick, c.Int("trunc"))
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

func picker(clips []string, value *string, truncate int) error {
	return huh.NewSelect[string]().
		Title("Pick a clip").
		OptionsFunc(func() []huh.Option[string] {
			opts := make([]huh.Option[string], len(clips))
			for i, clip := range clips {
				opts[i] = huh.NewOption(truncateString(clip, truncate), clip)
			}
			return opts
		}, nil).
		// Filtering(true).
		Value(value).
		WithTheme(huh.ThemeBase16()).
		Run()
}

func truncateString(s string, max int) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > max {
		return s[:max] + "..."
	}
	return s
}

func copyToClipboard(b []byte) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	<-clipboard.Write(clipboard.FmtText, b)
	return nil
}
