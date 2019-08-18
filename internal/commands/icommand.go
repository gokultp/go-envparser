package commands

// Command is the interface implemented by types
type Command interface {
	InitFlags()
	ParseFlags([]string)
	Help()
	Exec() error
}
