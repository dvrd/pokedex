package commands

import (
	"fmt"
)

type PokemonStat struct {
	BaseStat int    `json:"base_stat"`
	Stat     Entity `json:"stat"`
}

type PokemonType struct {
	Type Entity `json:"type"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

var PokedexMap map[string]Pokemon = map[string]Pokemon{}

func Pokedex(args []string) error {

	fmt.Println("Your Pokedex:")
	for _, val := range PokedexMap {
		fmt.Println("-", val)
	}

	return nil
}
