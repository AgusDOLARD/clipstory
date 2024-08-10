package commands

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AgusDOLARD/clipstory/socket"
	"github.com/urfave/cli/v2"
)

func (command *Command) Deamon() *cli.Command {
	return &cli.Command{
		Name:    "deamon",
		Aliases: []string{"d"},
		Usage:   "Starts the daemon",
		Action: func(c *cli.Context) error {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

			sock := socket.NewSocket(command.SocketPath)

			err := sock.Start(c.Context)
			if err != nil {
				return cli.Exit("failed to start daemon", 1)
			}

			<-done
			return sock.Stop()
		},
	}
}
