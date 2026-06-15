package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/christophrj/ocp-gen/logs"
)

const ociReplace = "// ocp-gen:replace"

var _ Command = &replaceCommand{}

// Activates on `//ocp-gen:replace`.
// Expects a variable number of search and replace pairs separated by '='.
// E.g. `//ocp-gen:replace search-a=ENV_A search-b=ENV_B`
// where ENV_* will be replaced by the result of the env variable lookup.
type replaceCommand struct {
	active    bool
	arguments []searchAndReplace
}

type searchAndReplace struct {
	search  string
	replace string
}

func NewReplaceCommand() Command {
	return &replaceCommand{}
}

// Execute implements [Command].
func (r *replaceCommand) Execute(loc string) string {
	if CommandPrefix(loc, ociReplace) {
		args := commandArguments(loc, ociReplace)
		r.arguments = []searchAndReplace{}
		for _, a := range args {
			spPair := strings.SplitN(a, "=", 2)
			if len(spPair) != 2 {
				logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid number of arguments", os.Getenv("GOFILE"), loc))
				return loc
			}
			replace, ok := os.LookupEnv(spPair[1])
			if !ok {
				logs.Debug(fmt.Sprintf("(%s) failed to lookup env (%s) of (%s)", os.Getenv("GOFILE"), spPair[1], loc))
			}
			r.arguments = append(r.arguments, searchAndReplace{search: spPair[0], replace: replace})
		}
		r.active = true
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the ocp-gen comment as part of the processing
		return ""
	}
	if r.active {
		original := loc
		for _, arg := range r.arguments {
			loc = strings.ReplaceAll(loc, arg.search, arg.replace)
		}
		logs.Debug(fmt.Sprintf("(%s) replaced (%s) with (%s)", os.Getenv("GOFILE"), original, loc))
		// replace is a one line command that instantly deactivates itself after processing a line of code
		r.active = false
	}
	return loc
}
