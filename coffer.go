package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/spaceshuttl/coffer/cmd"
)

// Command factories
var (
	NoteFactory *cli.CommandFactory
)

func main() {
	c := cli.NewCLI("Coffer", "2.0.0")
	c.Args = os.Args[1:]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands = map[string]cli.CommandFactory{
		"note": func() (cli.Command, error) {
			return cmd.NewNote(ui), nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
