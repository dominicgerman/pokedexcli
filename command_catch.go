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

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
  
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("did not catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught! Way to go sport!\n", pokemon.Name)
	
	return nil
}