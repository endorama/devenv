package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/endorama/devenv/command"
)

const (
	app = "devenv"
)

var (
	version = "0.1.0"
)

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{}

	commonUI := cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands["rehash"] = func() (cli.Command, error) {
		return &command.Rehash{
			UI: &commonUI,
		}, nil
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
