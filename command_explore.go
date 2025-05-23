package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)

	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}