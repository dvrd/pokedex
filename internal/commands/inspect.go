package commands

import (
	"fmt"
)

func Inspect(args []string) error {
	if len(args) < 1 {
		fmt.Println("Select a Pokemon to inspect")
		return nil
	}

	pokemon, found := PokedexMap[args[0]]

	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	if len(pokemon.Stats) > 0 {
		fmt.Println("Stats:")
	}
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}
	if len(pokemon.Types) > 0 {
		fmt.Println("Types:")
	}
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
