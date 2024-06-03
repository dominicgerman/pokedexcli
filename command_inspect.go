package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught that pokemon yet, you cheating fuck")
	}
  
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	
	return nil
}