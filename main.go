package main

import (
	"bufio"
	
	"fmt"
	
	"os"
	"strings"
	"time"

	"github.com/half-blood-prince-2710/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	Next     string
	Previous string
	Location string
}

var mp map[string]cliCommand
var cache *pokecache.Cache

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
			description: "print message for pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display names of location areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Go to previous location areas",
			callback:    commandMapBack,
		},
		"explore":{
			name:"explore",
			description :"List of all the pokemon located here",
			callback: commandExplore,
		},
	}
	cfg := &config{}

	cache = pokecache.NewCache(time.Second*10)
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		str := sc.Text()
		out := cleanInput(str)
		if len(out)==0 {
			fmt.Print("Please Enter Command.  Type 'help' to see available commands\n")
		}
		
		if cmd, exists := mp[out[0]]; exists {
			if cmd.name == "explore" {

				if len(out)>1{
					cfg.Location = out[1]
				} else {
					fmt.Print("Incorrect Command!. Please Type 'help' to see available commands\n")
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
