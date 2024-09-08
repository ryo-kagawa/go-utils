package commandline

type Command interface {
	Execute(arguments []string) (string, error)
}
