package main

import (
	"log"
	"os"

	"github.com/t10471/go-examples/grpc-examples/go/app/client"
	"github.com/t10471/go-examples/grpc-examples/go/app/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "client",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Value: "localhost:50051",
					},
				},
				Action: func(c *cli.Context) error {
					return client.ClientMain(
						c.String("address"),
					)
				},
			},
			{
				Name: "hw-simple",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "my-address",
						Value: "localhost:50051",
					},
					&cli.StringFlag{
						Name:  "other-address",
						Value: "localhost:50052",
					},
					&cli.BoolFlag{
						Name:  "unwrap-error",
						Value: true,
					},
				},
				Action: func(c *cli.Context) error {
					return server.HWSimpleMain(
						c.String("my-address"),
						c.String("other-address"),
						c.Bool("unwrap-error"),
					)
				},
			},
			{
				Name: "hw-complicated",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "my-address",
						Value: "localhost:50053",
					},
					&cli.StringFlag{
						Name:  "other-address",
						Value: "localhost:50052",
					},
				},
				Action: func(c *cli.Context) error {
					return server.HWComplicatedMain(
						c.String("my-address"),
						c.String("other-address"),
					)
				},
			},
			{
				Name: "other",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Value: "localhost:50053",
					},
					&cli.IntFlag{
						Name:  "sleep",
						Value: 3,
					},
				},
				Action: func(c *cli.Context) error {
					return server.OtherMain(
						c.String("address"),
						c.Int("sleep"),
					)
				},
			},
			{
				Name: "stream",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Value: "localhost:50054",
					},
				},
				Action: func(c *cli.Context) error {
					return server.StreamMain(
						c.String("address"),
					)
				},
			},
			{
				Name: "stream-client",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Value: "localhost:50054",
					},
					&cli.StringFlag{
						Name:  "method",
						Value: "server",
					},
				},
				Action: func(c *cli.Context) error {
					return client.StreamMain(
						c.String("address"),
						c.String("method"),
					)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
