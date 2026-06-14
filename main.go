package main

import (
	"fmt"
	"log"
	"os"

	"github.com/christophrj/ocp-gen/commands"
	"github.com/christophrj/ocp-gen/runner"
)

func main() {
	dryRun := commands.EvalBoolEnv("DRY_RUN")
	filepath := os.Getenv("PWD") + "/" + os.Getenv("GOFILE")

	log.Printf("generator called from file (%s) with the following flags...\n", filepath)
	log.Printf("\t-dry-run=%v", dryRun)

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

	// print in memory result
	_, _ = fmt.Fprint(os.Stdout, result.String())
}
