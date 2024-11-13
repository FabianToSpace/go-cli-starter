package main

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

var commands map[string]Command = make(map[string]Command)

func main() {
	example := newExampleCommand()
	commands[example.CliCommand.Name] = *example

	// If we don't provide any Arguments, we run the application using huh
	if len(os.Args) == 1 {
		if err := initialForm(commands); err != nil {
			log.Fatal(err)
		}
	} else {
		var flatCommands []*cli.Command
		for _, c := range commands {
			flatCommands = append(flatCommands, c.CliCommand)
		}

		app := &cli.App{
			Commands: flatCommands,
		}

		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
	}
}

func initialForm(commands map[string]Command) error {
	var action string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				OptionsFunc(func() []huh.Option[string] {
					var options []huh.Option[string]
					for _, command := range commands {
						options = append(options,
							huh.NewOption(command.CommandKey, command.CliCommand.Name))
					}
					return options
				}, nil).
				Value(&action),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if err := commands[action].FormFunc(); err != nil {
		return err
	}

	return nil
}
