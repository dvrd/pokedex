package commands

import (
	"github.com/dvrd/pokedex/internal/pokecache"
	"time"
)

type cliCommand struct {
	name        string
	description string
	Callback    func([]string) error
}

type Entity struct {
	Url  string
	Name string
}

type AreaResponse struct {
	Count    int
	Results  []Entity
	Next     string
	Previous string
}

var PokeCache *pokecache.Cache = pokecache.NewCache(10 * time.Second)
var baseLocationUrl string = "https://pokeapi.co/api/v2/location-area"
var basePokemonUrl string = "https://pokeapi.co/api/v2/pokemon"

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
		"explore": {
			name:        "explore",
			description: "Displays all pokemon on a given location",
			Callback:    Explore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			Callback:    Catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display details of acquired pokemon",
			Callback:    Inspect,
		},
	}
}
