package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type output struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func commandMap(cfg *config, _ string) error {
	if cfg.next == nil {
		fmt.Println("You're already on the last page")

		return nil
	}

	if val, ok := pokeAPICache.Get(*cfg.next); !ok {
		err := makeHttpRequestForMapCommand(*cfg.next, cfg)

		if err != nil {
			return err
		}
	} else {
		jsonOutput := output{}

		json.Unmarshal(val, &jsonOutput)

		for _, i := range jsonOutput.Results {
			fmt.Println(i.Name)
		}

		cfg.next = jsonOutput.Next
		cfg.previous = jsonOutput.Previous
	}

	return nil
}

func commandMapb(cfg *config, _ string) error {
	if cfg.previous == nil {
		fmt.Println("You're already on the first page")

		return nil
	}

	if val, ok := pokeAPICache.Get(*cfg.previous); !ok {
		err := makeHttpRequestForMapCommand(*cfg.previous, cfg)

		if err != nil {
			return err
		}
	} else {
		jsonOutput := output{}

		json.Unmarshal(val, &jsonOutput)

		for _, i := range jsonOutput.Results {
			fmt.Println(i.Name)
		}

		cfg.next = jsonOutput.Next
		cfg.previous = jsonOutput.Previous
	}

	return nil
}

func makeHttpRequestForMapCommand(url string, cfg *config) error {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}

	result, err := client.Do(request)

	if err != nil {
		return err
	}

	defer result.Body.Close()

	jsonOutput := output{}

	bytesOutput, err := io.ReadAll(result.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(bytesOutput, &jsonOutput)

	for _, i := range jsonOutput.Results {
		fmt.Println(i.Name)
	}

	pokeAPICache.Add(url, bytesOutput)

	cfg.next = jsonOutput.Next
	cfg.previous = jsonOutput.Previous

	return nil
}
