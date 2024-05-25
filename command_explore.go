package main

import (
    "fmt"
    "errors"
)

func callbackExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("no location area provided")
    }
    // location name is the first word of the input by user
    locationAreaName := args[0]
    locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
    if err != nil {
      return err 
    }
    fmt.Printf("Pokemon in %s:", locationArea.Name)
    for _, pokemon := range locationArea.PokemonEncounters{
        fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
    }
    return nil
}
