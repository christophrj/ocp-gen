package commands

import (
	"os"
	"strings"
)

func CommandPrefix(loc, commandIdentifier string) bool {
	return strings.HasPrefix(strings.TrimSpace(loc), commandIdentifier)
}

func commandArguments(loc, commandIdentifier string) []string {
	command := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(loc), commandIdentifier))
	return strings.Split(command, " ")
}

func EvalBoolEnv(envVar string) bool {
	v := strings.ToLower(os.Getenv(envVar))
	return v == "1" || v == "true"
}
