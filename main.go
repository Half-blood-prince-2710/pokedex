package main

import (
	"bufio"

	"fmt"

	"os"
	"strings"
	"time"

	"github.com/half-blood-prince-2710/pokedex/internal/pokecache"
)


var mp map[string]cliCommand
var cache *pokecache.Cache
var pokedex map[string]Pokemon

func cleanInput(text string) []string {

	list := strings.Fields(strings.ToLower(text))
	if len(list) > 0 {
		return list
	}
	return nil

}

func main() {
	mp = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print message for pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display names of location areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to previous location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List of all the pokemon located here. cmd:'explore location_name'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "To catch your favourite Pokemon. cmd:'catch pokemon_name'",
			callback:    commandCatch,
		},
		"inspect": {
			name:"inspect",
			description: "Details of Pokemon. cmd:'inspect pokemon_name'",
			callback: commandInspect,
		},
		"pokedex":{
			name :"pokedex",
			description: "List of caught Pokemon",
			callback: commandPokedex,
		},
	}
	pokedex := map[string]Pokemon{}
	cfg := &config{Pokedex: pokedex}

	cache = pokecache.NewCache(time.Second * 10)
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		str := sc.Text()
		out := cleanInput(str)
		if len(out) == 0 {
			fmt.Print("Please Enter Command.  Type 'help' to see available commands\n")
		}

		if cmd, exists := mp[out[0]]; exists {
			if cmd.name == "explore" || cmd.name == "catch" || cmd.name=="inspect" {

				if len(out) > 1 {
					if cmd.name == "explore" {
						cfg.Location = out[1]
					}
					if cmd.name == "catch" || cmd.name == "inspect" {
						cfg.Pokemon = out[1]
					}
					
				} else {
					fmt.Print("Incorrect Command!. Please Type 'help' to see available commands\n")
					continue
				}

			}
			if err := cmd.callback(cfg); err != nil {
						fmt.Print("Error: ", err, "\n")
					}

		} else {
			fmt.Print("Unknown Command. Type 'help' to see available commands\n")
		}
	}

}
