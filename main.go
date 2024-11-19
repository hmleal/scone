package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hmleal/scone/scoop"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "Scone - A high-performance replacement for Scoop",
		Description: "Scone - A high-performance replacement for Scoop Windows Instaler",
		Version:     "0.0.1",
		Compiled:    time.Now(),
		Commands: []*cli.Command{
			{
				Name:    "buckets",
				Aliases: []string{"s"},
				Usage:   "Manage Scone buckets.",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add a new bucket",
						Action: func(cCtx *cli.Context) error {
							scone, _ := scoop.NewScoop()
							scone.AddBucket(cCtx.Args().First())

							return nil
						},
					},
					{
						Name:  "rm",
						Usage: "Remove an existing bucket",
						Action: func(cCtx *cli.Context) error {
							scone, _ := scoop.NewScoop()
							scone.RemoveBucket(cCtx.Args().First())

							return nil
						},
					},
					{
						Name:  "known",
						Usage: "Show all known buckets you can use",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("NotImplementedError")

							return nil
						},
					},
					{
						Name:  "list",
						Usage: "List all added buckets",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("NotImplementedError")

							return nil
						},
					},
					{
						Name:  "update",
						Usage: "List all added buckets",
						Action: func(cCtx *cli.Context) error {
							scone, _ := scoop.NewScoop()
							scone.UpdateBuckets()

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
