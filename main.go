package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
}

var mp map[string]cliCommand
var cache *pokecache.Cache

func cleanInput(text string) string {

	list := strings.Fields(strings.ToLower(text))
	if len(list) > 0 {
		return list[0]
	}
	return ""

}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for key, val := range mp {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	return nil
}


func commandMap(cfg *config) error {
	var url string
	var resp []byte
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		url = cfg.Next
	}
	
	// Try getting the response from the cache
	resp, bool := cache.Get(url)
	if !bool { // If the response is not in cache, fetch it from the API
		fmt.Println("Fetching data from URL:", url) // Indicate fetching from URL
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error fetching data: %v", err)
		}
		defer res.Body.Close()

		resp, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}

		// Cache the fetched response
		cache.Add(url, resp)
	} else {
		fmt.Println("Fetching data from cache:", url) // Indicate fetching from cache
	}

	// Print the response for debugging
	// fmt.Printf("Response: %s\n", resp)

	// Define structures to parse the JSON response
	type Area struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type APIResponse struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []Area `json:"results"`
	}

	// Unmarshal the JSON response into the APIResponse struct
	var apiResponse APIResponse
	err := json.Unmarshal(resp, &apiResponse)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Update the Next and Previous fields in the config
	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous

	// Print the names of location areas
	for _, val := range apiResponse.Results {
		fmt.Println(val.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {

	var url string
	var resp []byte
	if cfg.Previous == "" {
		fmt.Print("you're on the first page\n")
		return nil
	} else {
		url = cfg.Previous
	}
	resp, bool := cache.Get(url)
	if bool == false {
		fmt.Println("Fetching data from URL:", url) 
		res, err := http.Get(url)
		// fmt.Print(res,"\n")
		if err != nil {
			return fmt.Errorf("error fetching data: %v", err)
		}
		defer res.Body.Close()

		resp, err = io.ReadAll(res.Body)
		// fmt.Print(resp,"\n")
		if err != nil {
			return fmt.Errorf("error coverting res : %v", err)
		}
		cache.Add(url, resp)
	}else {
		fmt.Println("Fetching data from cache:", url) // Indicate fetching from cache
	}

	type Area struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type APIResponse struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []Area `json:"results"`
	}

	var apiResponse APIResponse

	err := json.Unmarshal(resp, &apiResponse)

	
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous
	// fmt.Print(apiResponse.Previous,"\n")
	for _, val := range apiResponse.Results {
		fmt.Print(val.Name, "\n")
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
	}
	cfg := &config{}

	cache = pokecache.NewCache(time.Second*10)
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		str := sc.Text()
		out := cleanInput(str)
		if cmd, exists := mp[out]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Print("Error: ", err, "\n")
			}
		} else {
			fmt.Print("Unknown Command. Type 'help' to see available commands\n")
		}
	}

}
