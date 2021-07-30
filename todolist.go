package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"todo_cli/ado"
	"todo_cli/services"

	"github.com/urfave/cli/v2"
)

func main() {

	ado.Initialize() //initialize de database

	app := &cli.App{
		Name:  "todolist",
		Usage: "Todo list on simple cli",
		Action: func(c *cli.Context) error {
			fmt.Println(" *** TODO List ***")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add new task to list",
				Action: func(c *cli.Context) error {
					services.Add(c.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "Complete task with id parameter",
				Action: func(c *cli.Context) error {
					taskId, err := strconv.Atoi(c.Args().First())

					if err != nil {
						log.Fatal(err)
					}

					services.Complete(taskId)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List of tasks on todo",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Usage:   "List all include completed tasks",
						Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					listAll := false

					listAll = c.Bool("all")

					services.List(listAll)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
