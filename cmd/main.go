package main

import (
	"log"
	"os"

	commands "github.com/AdhityaRamadhanus/thug.ly/commands"
	cli "github.com/urfave/cli"
)

func cmdNotFound(c *cli.Context, command string) {
	log.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}

func main() {
	app := cli.NewApp()
	app.Name = "thugly"
	app.Author = "Adhitya Ramadhanus"
	app.Email = "adhitya.ramadhanus@gmail.com"

	app.Action = commands.Thuglify
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input",
			Value: "",
			Usage: "Image input path",
		},
		cli.StringFlag{
			Name:  "text",
			Value: "",
			Usage: "Text to add at the bottom",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "thuglify.jpg",
			Usage: "Image output path",
		},
	}
	app.CommandNotFound = cmdNotFound // Inspired by docker machine
	app.Usage = "Thuglify your photos"
	app.Version = "1.0.0"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
