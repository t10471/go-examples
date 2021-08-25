package main

import (
	"log"
	"os"

	"github.com/t10471/go-examples/csv-examples/app/client"
	"github.com/t10471/go-examples/csv-examples/app/csv"
	"github.com/t10471/go-examples/csv-examples/app/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "host",
						Value:   "localhost:8099",
					},
					&cli.StringFlag{
						Name:    "csv",
						Value:   "large.csv",
					},
					&cli.BoolFlag{
						Name:    "add-content-length",
						Value:   false,
					},
				},
				Action:  func(c *cli.Context) error {
					return server.HttpMain(
						c.String("host"),
						c.String("csv"),
						c.Bool("add-content-length"),
						)
				},
			},
			{
				Name:    "client",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "host",
						Value:   "localhost:8099",
					},
					&cli.StringFlag{
						Name:    "csv",
						Value:   "large.csv",
					},
					&cli.BoolFlag{
						Name:    "use-read-all",
						Value:   false,
					},
				},
				Action:  func(c *cli.Context) error {
					return client.HttpMain(
						c.String("host"),
						c.String("csv"),
						c.Bool("use-read-all"))
				},
			},
			{
				Name:    "gen-csv",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "line",
						Value:   100000,
					},
					&cli.StringFlag{
						Name:    "csv",
						Value:   "large.csv",
					},
				},
				Action:  func(c *cli.Context) error {
					return csv.MakerMain(c.Int("line"), c.String("name"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}