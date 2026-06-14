package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/christophrj/ocp-gen/commands"
)

const ocpgen = "//go:generate ocp-gen"

// Runner takes a set of Commands
type Runner struct {
	Commands []commands.Command
}

// Run executes the runner commands on the current go generate file
func (r *Runner) Run(fpath string) (result bytes.Buffer) {
	// iterate over each file
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		original := fileScanner.Text()
		loc := original
		for _, c := range r.Commands {
			loc = c.Execute(loc)
		}
		if commands.CommandPrefix(loc, ocpgen) {
			// remove go:generate ocp-gen lines
			loc = ""
		}
		// add newline unless line should be removed
		if original == loc || loc != "" {
			if _, err := fmt.Fprintln(&result, loc); err != nil {
				panic(err)
			}
		}
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return result
}
