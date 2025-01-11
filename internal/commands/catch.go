package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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

var Pokedex map[string]Pokemon = map[string]Pokemon{}

func Catch(args []string) error {
	if len(args) < 1 {
		fmt.Println("Select a Pokemon to attempt to catch")
		return nil
	}

	url := fmt.Sprintf("%s/%s", basePokemonUrl, args[0])
	bodyBytes, found := PokeCache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("request failed with status %v", res.Status)
		}

		bodyBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		PokeCache.Add(url, bodyBytes)
	}

	var jsonResponse Pokemon
	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	err := decoder.Decode(&jsonResponse)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", jsonResponse.Name)
	random := rand.Intn(jsonResponse.BaseExperience)
	if ((random * 100) / jsonResponse.BaseExperience) >= 50 {
		fmt.Printf("%s was caught!\n", jsonResponse.Name)
		Pokedex[jsonResponse.Name] = jsonResponse
	} else {
		fmt.Printf("%s escaped!\n", jsonResponse.Name)
	}

	return nil
}
