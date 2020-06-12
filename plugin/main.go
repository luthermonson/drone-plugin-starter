package plugin

import (
	"fmt"

	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

/*
example YAML step settings
settings:
  username: jack
  files:
   - file1
   - file2
*/

var Flags = []cli.Flag{
	/* add your Settings flags
	cli.StringFlag{
	    Name:  "username",
	    EnvVars: []string{"PLUGIN_USERNAME"},
	},
	cli.StringSliceFlag{
	    Name:  "files",
	    EnvVars: []string{"PLUGIN_FILES"},
	},
	*/
}

type Settings struct {
	/* add your settings
	username string
	files []string
	*/
}

func NewSettingsFromContext(c *cli.Context) Settings {
	return Settings{
		/* map c.String, c.StringSlice, c.Bool to your config fields
		username: c.String("username"),
		username: c.StringSlice("files"),
		*/
	}
}

func Exec(pipeline drone.Pipeline, settings Settings) error {
	// plugin logic goes here
	fmt.Println(pipeline.Repo.Name)
	// use config p.Config

	return nil
}
