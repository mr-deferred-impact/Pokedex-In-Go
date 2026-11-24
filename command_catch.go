package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

type pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

var pokedex map[string]pokemon = map[string]pokemon{}

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Please enter the name of the pokemon to catch!")
	}

	err := makeHTTPRequestForCommandCatch(pokemonName)

	if err != nil {
		return err
	}

	return nil
}

func makeHTTPRequestForCommandCatch(pokemonName string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Please enter a valid pokemon name!")
	}

	defer response.Body.Close()

	myPokemon := pokemon{}
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&myPokemon)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const RANDOM_THRESHOLD int = 35

	baseExperience := myPokemon.BaseExperience
	caught := rand.Intn(baseExperience)

	if caught <= RANDOM_THRESHOLD {
		pokedex[myPokemon.Name] = myPokemon
		fmt.Println(myPokemon.Name, "was caught!")
	} else {
		fmt.Println(myPokemon.Name, "escaped!")
	}

	return nil
}
