package commandline

type SubCommand interface {
	Command
	Name() string
}
