package main


type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	Next     string
	Previous string
	Location string
	Pokemon  string
	Pokedex  map[string]Pokemon
}
type Pokemon struct {
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	BaseExperience int `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
