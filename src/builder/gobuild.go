// Example usage
//. .\gobuild.exe -file=build.go compile test jsmin -params={version:4.0}

package main

import (
	"builder/models"
	"encoding/json"
	"github.com/codegangsta/cli"
	"os"
)

var file string
var params string
var parameters map[string]interface{}
var tasks []string

func main() {
	app := cli.NewApp()
	app.Name = "go build"
	app.Usage = "build multiple step operations using tasks"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "show more output",
		},
		cli.StringFlag{
			Name:  "f, file",
			Usage: "Specify a build file.",
		},
		cli.StringFlag{
			Name:  "p, params",
			Usage: "The parameters passed to the task",
		},
		cli.StringFlag{
			Name:  "t, tasks",
			Usage: "The tasks to be run",
		},
	}

	app.Action = Build

	app.Run(os.Args)

	gobuild := new(models.Runner)
	gobuild.Invoke(file, tasks, parameters)
}

func Build(c *cli.Context) {

	if len(c.Args()) > 0 {
		tasks = c.Args()[0]
	}

	if len(c.String("params")) > 0 {
		if len(params) > 0 {
			paramByte := []byte(params)
			if err := json.Unmarshal(paramByte, &parameters); err != nil {
				panic(err)
			}
		}
	}
}
