package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ritarock/athenaq/lib/action"
	athenaqfile "github.com/ritarock/athenaq/lib/file"

	"github.com/urfave/cli/v2"
)

func main() {
	var profile string
	var bucket string
	var key string
	var query string = ""
	var file string = ""

	app := cli.App{
		Name:  "athenaq",
		Usage: "Run athena query",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "profile",
				Usage:       "set aws profile",
				Destination: &profile,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "bucket",
				Usage:       "set aws bucket",
				Destination: &bucket,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "key",
				Usage:       "set aws bucket key",
				Destination: &key,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "query",
				Usage:       "set athena query",
				Destination: &query,
			},
			&cli.StringFlag{
				Name:        "file",
				Usage:       "set file athena query",
				Destination: &file,
			},
		},
		Action: func(c *cli.Context) error {
			if query == "" && file == "" {
				fmt.Println("No query is set.")
			} else if file != "" {
				query = athenaqfile.Read(file)
			}
			action.Run(profile, bucket, key, query)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
