package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func callbackCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("no pokemon name  provided")
    }
    // location name is the first word of the input by user
    pokemonName := args[0]
    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
      return err 
    }
    const threshold = 50
    randNum := rand.Intn(pokemon.BaseExperience)
    fmt.Printf("Throwing pokeball at %s...\n", pokemonName)
    if randNum > threshold {
        return fmt.Errorf("%s ran away!", pokemonName)
    }
    cfg.caughtPokemon[pokemonName] = pokemon
    fmt.Printf("%s was caught\n", pokemonName)
    return nil
}

