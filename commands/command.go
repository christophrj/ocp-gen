package commands

type Command interface {
	Execute(loc string) string
}
