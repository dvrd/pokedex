package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var next int = 0
var prev int = 0
var limit int = 20

func getArea(offset int, dir string) ([]Entity, error) {
	url := fmt.Sprintf("%s?offset=%d&limit=%d", baseLocationUrl, offset, limit)
	bodyBytes, found := PokeCache.Get(url)

	if !found {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("request failed with status %v", res.Status)
		}

		bodyBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		PokeCache.Add(url, bodyBytes)
	}

	var jsonResponse AreaResponse
	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	err := decoder.Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	if dir == "prev" && prev >= 20 {
		prev -= limit
	} else if jsonResponse.Count-offset > 20 {
		prev = offset
	}

	if dir == "next" && next <= jsonResponse.Count-20 {
		next += limit
	} else if next > 0 {
		next = prev
	}

	return jsonResponse.Results, nil
}

func Map(args []string) error {
	locations, err := getArea(next, "next")
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area.Name)
	}

	return nil
}

func MapPrevious(args []string) error {
	if prev == 0 && next == 0 {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, err := getArea(prev, "prev")
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area.Name)
	}

	return nil
}
