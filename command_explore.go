package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type commandExploreOutput struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, locationArea string) error {
	if locationArea == "" {
		return errors.New("location area name not entered.")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationArea)

	if val, ok := pokeAPICache.Get(url); ok {
		unmarshalJsonDataAndPrintToConsoleForCommandExplore(val, locationArea)
	} else {
		err := makeHTTPRequestForExploreCommand(locationArea)

		if err != nil {
			return err
		}
	}

	return nil
}

func makeHTTPRequestForExploreCommand(locationArea string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationArea)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("The location area you specified doesn't exist. Please enter a valid location area name!")
	}

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = unmarshalJsonDataAndPrintToConsoleForCommandExplore(bodyBytes, locationArea)

	if err != nil {
		return err
	}

	pokeAPICache.Add(url, bodyBytes)

	return nil
}

func unmarshalJsonDataAndPrintToConsoleForCommandExplore(dataBytes []byte, locationArea string) error {
	myOutput := commandExploreOutput{}

	err := json.Unmarshal(dataBytes, &myOutput)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")

	for _, data := range myOutput.PokemonEncounters {
		fmt.Println(" -", data.Pokemon.Name)
	}

	return nil
}
