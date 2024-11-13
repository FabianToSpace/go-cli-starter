package main

import "github.com/urfave/cli/v2"

type Command struct {
	CommandKey string
	FormFunc   func() error
	CliCommand *cli.Command
}

func newCommand(commandKey string, formFunc func() error, cliCommand *cli.Command) *Command {
	return &Command{
		CommandKey: commandKey,
		FormFunc:   formFunc,
		CliCommand: cliCommand,
	}
}
