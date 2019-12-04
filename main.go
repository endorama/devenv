package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/endorama/devenv/command"
	"github.com/endorama/devenv/internal/version"
)

const (
	app = "devenv"
)

func main() {
	c := cli.NewCLI(app, version.Version())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{}

	commonUI := cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands["list"] = func() (cli.Command, error) {
		return &command.List{
			UI: &commonUI,
		}, nil
	}
	c.Commands["rehash"] = func() (cli.Command, error) {
		return &command.Rehash{
			UI: &commonUI,
		}, nil
	}
	c.Commands["shell"] = func() (cli.Command, error) {
		return &command.Shell{
			UI: &commonUI,
		}, nil
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
