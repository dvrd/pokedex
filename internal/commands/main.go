package commands

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func Get() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"map": {
			name:        "map",
			description: "Displays name of next 20 location",
			Callback:    Map,
		},
		"mapb": {
			name:        "map",
			description: "Displays name of previous 20 location",
			Callback:    MapPrevious,
		},
	}
}
