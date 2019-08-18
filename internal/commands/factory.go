package commands

var (
	genCmd     = NewGenerate()
	helpCmd    = NewHelp()
	versionCmd = NewVersion()
)

// initializes when package initializes
func init() {
	genCmd.InitFlags()
	helpCmd.InitFlags()
	versionCmd.InitFlags()
}

// GetCmd will get command by flags
func GetCmd(args []string) Command {
	switch args[0] {
	case "generate":
		genCmd.ParseFlags(args[1:])
		return genCmd
	case "version":
		return versionCmd
	default:
		return helpCmd
	}
}
