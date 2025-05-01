package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return  errors.New("you must provide a pokemon to inspect")
	}

	name := args[0]
	val, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you didn't catch that pokemon yet")
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range val.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeInfo := range val.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}