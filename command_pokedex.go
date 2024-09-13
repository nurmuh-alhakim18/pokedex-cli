package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
	}

	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
