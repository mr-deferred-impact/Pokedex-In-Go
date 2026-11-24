package main

import "fmt"

func commandHelp(cfg *config, _ string) error {
	helpCommandOutput := `Welcome to the Pokedex!

Usage:

`

	for key := range cfg.supportedCommands {
		helpCommandOutput += fmt.Sprintf("%s: %s\n", cfg.supportedCommands[key].name, cfg.supportedCommands[key].description)
	}

	fmt.Println(helpCommandOutput)

	return nil
}
