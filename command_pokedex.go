package main
import (
    "fmt"
)

func callbackPokedex(cfg *config, args... string) error {
    caughtPokemon := cfg.caughtPokemon
    fmt.Println("Your Pokemon:")
    for _, pokemon := range caughtPokemon {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    return nil
}

