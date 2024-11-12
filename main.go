package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Command string
	Args    []string
}

var command Command

func main() {
	// If we don't provide any Arguments, we run the application using huh
	if len(os.Args) == 1 {
		if err := initialForm(); err != nil {
			log.Fatal(err)
		}
	} else {
		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name:    "cmd1",
					Aliases: []string{"c"},
					Usage:   "Use cmd1",
					Action: func(cCtx *cli.Context) error {
						command.Command = cCtx.Command.Name
						command.Args = mergeArgs(cCtx.Args())
						return nil
					},
				},
			},
		}

		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Executing command %v\n", command)
}

func mergeArgs(args cli.Args) []string {
	var argsMerged []string

	for i := range args.Len() {
		argsMerged = append(argsMerged, args.Get(i))
	}

	return argsMerged
}

func initialForm() error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(
					huh.NewOption("Option 1", "op1"),
					huh.NewOption("Option 2", "op2"),
				).
				Value(&command.Command),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	return nil
}
