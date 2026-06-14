package commands

import (
	"os"
	"strings"
	"unicode"
)

func CommandPrefix(loc, commandIdentifier string) bool {
	return strings.HasPrefix(strings.TrimLeftFunc(loc, unicode.IsSpace), commandIdentifier)
}

func commandArguments(loc, commandIdentifier string) []string {
	trimPrefix := strings.TrimLeftFunc(strings.TrimPrefix(strings.TrimLeftFunc(loc, unicode.IsSpace), commandIdentifier), unicode.IsSpace)
	return strings.Split(trimPrefix, " ")
}

func EvalBoolEnv(envVar string) bool {
	v := strings.ToLower(os.Getenv(envVar))
	return v == "1" || v == "true"
}
