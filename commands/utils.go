package commands

import (
	"os"
	"strings"
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

func EvalBoolEnv(envVar string) bool {
	v := strings.ToLower(os.Getenv(envVar))
	return v == "1" || v == "true"
}
