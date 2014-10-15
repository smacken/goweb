// Example usage
//. .\gobuild.exe -file=build.go compile test jsmin -params={version:4.0}

package main

import (
	"encoding/json"
	"flag"
	"os"
	"runner"
)

var file string
var params string

func main() {
	tasks := os.Args[1:]
	flag.StringVar(&file, "file", "default.go", "The task file to run. (default.go by default)")
	flag.StringVar(&params, "params", "", "The parameters passed to the tasks.")

	flag.Parse()

	var parameters map[string]interface{}
	if len(params) > 0 {
		if err := json.Unmarshal(&params, &parameters); err != nil {
			panic(err)
		}
	}

	gobuild := new(runner.Runner)
	gobuild.Invoke(file, tasks, parameters)
}
