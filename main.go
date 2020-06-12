package main

import (
	"log"
	"os"

	urfave "github.com/drone-plugins/drone-plugin-lib/urfave"
	"github.com/drone/drone-plugin-starter/plugin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

var version string // build number set at compile-time

func main() {
	app := &cli.App{
		Name:    "my plugin",
		Usage:   "my plugin usage",
		Action:  run,
		Version: version,
		Flags:   append(urfave.Flags(), plugin.Flags...),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	return plugin.Exec(
		urfave.PipelineFromContext(c),
		plugin.NewSettingsFromContext(c),
	)
}
