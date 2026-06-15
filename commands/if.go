package commands

import (
	"fmt"
	"os"

	"github.com/christophrj/ocp-gen/logs"
)

const (
	ocpIf = "// ocp-gen:if"
	ocpFi = "// ocp-gen:fi"
)

var _ Command = &ifCommand{}

// Activates on `//ocp-gen:if`.
// Expects an env variable as parameter that holds a bool value.
// e.g. `//ocp-gen:if ENV_FEATURE
// where ENV_FEATURE = 'false' will result in any following line to be removed
// until `//ocp-gen-fi` deactivates the command again.
type ifCommand struct {
	active    bool
	condition bool
}

func NewIfCommand() Command {
	return &ifCommand{}
}

// Execute implements [Command].
func (r *ifCommand) Execute(loc string) string {
	if CommandPrefix(loc, ocpIf) {
		args := commandArguments(loc, ocpIf)
		if len(args) > 1 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid number of arguments", os.Getenv("GOFILE"), loc))
		}
		r.active = true
		r.condition = EvalBoolEnv(args[0])
		logs.Debug(fmt.Sprintf("ifCommand condition = %v", r.condition))
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the ocp-gen comment as part of the processing
		return ""
	}
	if CommandPrefix(loc, ocpFi) {
		r.active = false
		r.condition = false
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the ocp-gen comment as part of the processing
		return ""
	}
	if r.active && !r.condition {
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		return ""
	}
	return loc
}
