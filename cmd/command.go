package cmd

import (
	"einfach-msg/cmd/served"
	"os"

	log "github.com/sirupsen/logrus"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Execute() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Warn(`please specify your command, available command is: serve`)
		return
	}

	// TODO create function to load config
	serve := served.New()

	switch args[0] {
	case "serve":
		serve.HTTP()
	default:
		log.Info(`command not available, available command is: serve`)
	}
}
