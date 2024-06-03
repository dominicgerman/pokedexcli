package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	fmt.Println("You done caught these pokefucks:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("Name: %v\n", pokemon.Name)
	}
	
	return nil
}