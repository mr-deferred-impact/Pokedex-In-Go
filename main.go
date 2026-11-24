package main

import (
	"time"

	"github.com/mr-deferred-impact/pokedexcli/internal/pokecache"
)

var pokeAPICache pokecache.Cache = pokecache.NewCache(5 * time.Minute)

func main() {
	cfg := config{}
	initialUrl := "https://pokeapi.co/api/v2/location-area/"

	cfg.next = &initialUrl

	cfg.supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Prints this help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Paginate forward by 20 results for the location areas",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Paginate backward by 20 results for the location areas",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore",
			description: "Explore location area by name or ID.",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon by name or ID.",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect",
			description: "Return the stats and information about a caught pokemon.",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Return the name of the pokemons in your pokedex.",
			callback:    commandPokedex,
		},
	}

	startRepl(&cfg)
}
