package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Encounter struct {
	Pokemon Entity `json:"pokemon"`
}

type EncountersResponse struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

func Explore(args []string) error {
	url := fmt.Sprintf("%s/%s", baseLocationUrl, args[0])
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

	var jsonResponse EncountersResponse
	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	err := decoder.Decode(&jsonResponse)
	if err != nil {
		return err
	}

	for _, encounter := range jsonResponse.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
