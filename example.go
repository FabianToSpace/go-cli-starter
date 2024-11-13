package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

type ExampleCommand struct {
	Command
}

func newExampleCommand() *Command {
	var command ExampleCommand

	command.Command = *newCommand("Example Command", command.initForm, command.initCliCommand())

	return &command.Command
}

func (e *ExampleCommand) initForm() error {
	var option string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Example Command Title").
				Options(
					huh.NewOption("Test Option", "test"),
					huh.NewOption("Test Option 2", "test2"),
				).
				Value(&option),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Selected Option: %s\n", option)

	return nil
}

func (e *ExampleCommand) initCliCommand() *cli.Command {
	return &cli.Command{
		Name:    "example",
		Aliases: []string{"e"},
		Usage:   "example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "option",
				Usage: "[test|test2]",
				Value: "test",
			},
		},
		Action: func(cCtx *cli.Context) error {
			option := cCtx.String("option")
			fmt.Printf("Selected Option: %s\n", option)

			return nil
		},
	}
}
