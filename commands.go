package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	
)

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
	resp, found := cache.Get(url)
	if !found { // If the response is not in cache, fetch it from the API
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
	resp, found := cache.Get(url)
	if found == false {
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

func commandExplore(cfg *config) error {
	var url string
	var resp []byte
	url = fmt.Sprint("https://pokeapi.co/api/v2/location-area/",cfg.Location)
	resp,found :=cache.Get(url)
	if !found {
		
		res,err := http.Get(url)
		
		if err!=nil {
			fmt.Print("Error fetching data: ",err)
			return nil
		}
		defer res.Body.Close()
		if res.StatusCode == 404 {
			fmt.Print("Warning! Warning! Warning! Please Enter a valid Location Name !, Otherwise I will Kick you XD\n")
			return nil
		}
		resp,err=io.ReadAll(res.Body)
		cache.Add(url,resp)
	}
	// fmt.Print("res",resp,"\n\n")
	type Pokemon struct {
		Name string `json:"name"`
	}
	
	var PokemonList struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

	err:=json.Unmarshal(resp,&PokemonList)
	if err!=nil {
		
		fmt.Print("Error Unmarshaling data : ",err,"\n")
	}
	for _,val := range PokemonList.PokemonEncounters {
		fmt.Print(val.Pokemon.Name,"\n")
	}


	return nil
}

func commandCatch(cfg *config) error {
	// rand.Seed(time.Now().UnixNano())
	fmt.Print("Throwing a Pokeball at  ",cfg.Pokemon,"...\n")
	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/",cfg.Pokemon,"/")
	// fmt.Print("url ",url ,"\n")
	resp,found :=cache.Get(url)
	if !found {
		
		res,err := http.Get(url)
		
		if err!=nil {
			fmt.Print("Error fetching data: ",err)
			return nil
		}
		defer res.Body.Close()
		if res.StatusCode == 404 {
			fmt.Print("Warning! Warning! Warning! Please Enter a valid Pokemon Name !, Otherwise I will caught you XD\n")
			return nil
		}
		resp,err=io.ReadAll(res.Body)
		cache.Add(url,resp)
	}
	var input struct {
		Name string `json="name"`
		Base_Experience int `json="base_experience"`
	}
	err:=json.Unmarshal(resp,&input)
	if err!=nil {
		
		fmt.Print("Error Unmarshaling data : ",err,"\n")
	}
	
	val := rand.Intn(2*input.Base_Experience)
	fmt.Print(input.Base_Experience," name: ",input.Name,"rand: ",val,"\n")
	if  val > input.Base_Experience {
		fmt.Print(cfg.Pokemon," was caught!\n")
		cfg.Pokedex[url] = Pokemon{Name: cfg.Pokemon}
	}else {
		fmt.Print(cfg.Pokemon," escaped!\n")
	}
	return nil
}