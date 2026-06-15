package main

import (
	"fmt"
	"log"
	"os"

	"github.com/christophrj/ocp-gen/commands"
	"github.com/christophrj/ocp-gen/logs"
	"github.com/christophrj/ocp-gen/runner"
)

func main() {
	dryRun := commands.EvalBoolEnv("DRY_RUN")
	debug := commands.EvalBoolEnv("DEBUG")

	logs.Init(debug)

	filepath := os.Getenv("PWD") + "/" + os.Getenv("GOFILE")
	logs.Debug(fmt.Sprintf("generator called from file (%s) with the following flags...", filepath))

	runner := runner.Runner{
		Commands: []commands.Command{
			commands.NewReplaceCommand(),
			commands.NewIfCommand(),
		},
	}
	result := runner.Run(filepath)

	if !dryRun {
		if err := os.WriteFile(filepath, result.Bytes(), 0644); err != nil {
			panic(err)
		}
		log.Printf("%s: saved changes\n", filepath)
		return
	}

	// dry run prints in memory result unless debug is set
	if !debug {
		fmt.Fprintf(os.Stdout, "### %s\n", filepath)
		_, _ = fmt.Fprint(os.Stdout, result.String())
	}
}
