package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("this command doesn't take arguments")
	}

	fmt.Printf("Your Pokedex:\n")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", cfg.caughtPokemon[pokemon].Name)
	}

	return nil
}