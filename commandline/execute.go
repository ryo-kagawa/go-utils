package commandline

func Execute(rootCommand RootCommand, arguments []string, subCommands ...SubCommand) (string, error) {
	if 1 <= len(arguments) {
		argument := arguments[0]
		for _, subCommand := range subCommands {
			if argument == subCommand.Name() {
				return subCommand.Execute((arguments[1:]))
			}
		}
	}
	return rootCommand.Execute(arguments)
}

func Find(rootCommand RootCommand, arguments []string, subCommands ...SubCommand) (Command, []string) {
	if 1 <= len(arguments) {
		argument := arguments[0]
		for _, subCommand := range subCommands {
			if argument == subCommand.Name() {
				return subCommand, arguments[1:]
			}
		}
	}
	return rootCommand, arguments
}
