package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        
        fmt.Println("Pokedex >")
        scanner.Scan()
        text := scanner.Text()
        cleaned := cleanInput(text)

        if len(cleaned) == 0 {
            continue
        }
        commandName := cleaned[0]
        args := []string{}
        if len(cleaned) > 1 {
            args = cleaned[1:]
        }

        availableCommands := getCommands()
        command, ok := availableCommands[commandName]
        if !ok {
            fmt.Println("invalid command")
            continue
        }
        err := command.callback(cfg, args...)
        if err != nil {
            fmt.Println(err)
        }
    }
}



func cleanInput(str string) []string {
    lowered := strings.ToLower(str)
    words := strings.Fields(lowered)
    return words 
}
    
type cliCommand struct {
    name string
    description string
    callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Display help message",
            callback:    callbackHelp,
        },
        "map": {
            name:   "map",
            description: "Display name of locations in the Pokemonworld",
            callback: callbackMap,
        },
        "mapb": {
            name: "mapb",
            description: "Display the 20 previous locations",
            callback: callbackMapb,
        },
        "explore": {
            name: "explore {location area}",
            description: "List pokemon in location area",
            callback: callbackExplore,
        },
        "catch": {
            name:        "catch {pokemon}",
            description: "Attempt to catch pokemon and add it to your pokedex",
            callback:     callbackCatch,
        },
        "inspect": {
            name:        "inspect {pokemon}",
            description: "Shows information about pokemon you have caprtured",
            callback:     callbackInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "Shows all the pokemon that you have caught",
            callback:     callbackPokedex,
        },
        "exit": {
            name:        "exit",
            description: "Exit the pokedex",
            callback:     callbackExit,
        },
    }
}
