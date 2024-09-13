package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 20
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", resp.Name)
	}

	cfg.pokedex[pokemonName] = resp
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
