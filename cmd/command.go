package cmd

import (
	"einfach-msg/cmd/served"
	"os"

	log "github.com/sirupsen/logrus"
)

// Command is a main command which will instruct all of the available service
type Command struct{}

// New will instantiate Command itself
func New() *Command {
	return &Command{}
}

// Execute will read argument and run the command if it's valid
func (c *Command) Execute() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Warn(`please specify your command, available command is: serve`)
		return
	}

	serve := served.New()

	switch args[0] {
	case "serve":
		serve.HTTP()
	default:
		log.Info(`command not available, available command is: serve`)
	}
}
