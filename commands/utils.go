package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/christophrj/ocp-gen/logs"
)

func CommandPrefix(loc, commandIdentifier string) bool {
	line := strings.TrimSpace(loc)
	if !strings.HasPrefix(line, "//") {
		return false
	}
	uncommentedCommand := uncommentLine(line)
	return strings.HasPrefix(uncommentedCommand, commandIdentifier)
}

func uncommentLine(line string) string {
	return strings.TrimSpace(strings.TrimPrefix(line, "//"))
}

func trimCommand(line, commandIdentifier string) string {
	return strings.TrimSpace(strings.TrimPrefix(uncommentLine(strings.TrimSpace(line)), commandIdentifier))
}

func commandArguments(loc, commandIdentifier string) []string {
	args := trimCommand(loc, commandIdentifier)
	return strings.Split(args, " ")
}

func assignments(loc, commandIdentifier string) []assignment {
	args := commandArguments(loc, commandIdentifier)
	assignments := []assignment{}
	for _, a := range args {
		pair := strings.SplitN(a, "=", 2)
		if len(pair) != 2 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid argument assignment", os.Getenv("GOFILE"), loc))
			return nil
		}
		assignments = append(assignments, assignment{left: pair[0], right: pair[1]})
	}
	return assignments
}

func EvalBoolEnv(envVar string) bool {
	v := strings.ToLower(os.Getenv(envVar))
	return v == "1" || v == "true"
}

type assignment struct {
	left  string
	right string
}
