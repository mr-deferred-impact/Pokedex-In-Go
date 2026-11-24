package main

import "fmt"

func commandPokedex(*config, string) error {
	fmt.Println("Your Pokedex:")

	for i := range pokedex {
		fmt.Println(" -", i)
	}

	return nil
}
