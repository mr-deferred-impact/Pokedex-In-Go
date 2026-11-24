package main

import "fmt"

func commandInspect(_ *config, pokemonName string) error {
	if val, ok := pokedex[pokemonName]; ok {
		fmt.Println("Name:", val.Name)
		fmt.Println("Height:", val.Height)
		fmt.Println("Weight:", val.Weight)
		fmt.Println("Stats:")

		for _, i := range val.Stats {
			fmt.Printf(" - %s: %d\n", i.Stat.Name, i.BaseStat)
		}

		fmt.Println("Types:")

		for _, i := range val.Types {
			fmt.Printf(" -%s\n", i.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
