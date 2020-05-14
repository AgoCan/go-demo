package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func example(c *cli.Context) error {
	name := "hahaha"
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}
	if c.String("lang") == "spanish" {
		fmt.Println("spanish", name)
	} else {
		fmt.Println("Hello", name)
	}
	return nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Value:   "english",
				Aliases: []string{"l"},
				Usage:   "language for the greeting",
			},
		},
		Action: example,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
