package cli

var commands = map[string]Command{}

func Register(name string, c Command) {
	commands[name] = c
}

func Commands() map[string]Command {
	return commands
}
