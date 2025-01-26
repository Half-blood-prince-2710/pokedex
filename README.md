# Pokédex CLI

A **Pokédex CLI** is a command-line tool written in Go, designed to help you explore the world of Pokémon. With this tool, you can search for Pokémon details, check Pokémon availability in specific locations, catch Pokémon, and manage your collection—all from your terminal!

## Features
- 🐾 **Browse Pokémon**: Search for Pokémon by name or ID and view their stats, abilities, and types.  
- 🌍 **Location-based Search**: Find Pokémon available in specific locations.  
- 🎣 **Catch Pokémon**: Attempt to catch Pokémon and add them to your collection.  
- 🕐 **Inspect Caught Pokémon**: View details of Pokémon you've caught, including stats and abilities.  
- 📂 **View Caught Pokémon**: Display all Pokémon in your collection.  
- ⚡ **Caching for Performance**: Speeds up repeated requests with locally cached data.  

## Tech Stack
- **Language**: Go  
- **API**: [PokéAPI](https://pokeapi.co/)  

## Installation

### Prerequisites
- Go (1.18 or later) installed on your machine  

### Steps
1. Clone the repository:  
   ```bash  
   git clone https://github.com/half-blood-prince-2710/pokedex-cli.git  
   cd pokedex-cli  
   ```  

2. Install dependencies:  
   ```bash  
   go mod tidy  
   ```  

3. Build the CLI tool:  
   ```bash  
   go build -o pokedex  
   ```  

4. Run the CLI tool:  
   ```bash  
   ./pokedex  
   ```  

## Usage
### Commands
- **help**: Print usage information for the Pokédex.  
  ```bash  
  ./pokedex help  
  ```
  Output:
  ```
  Welcome to the Pokedex!
  Usage:
  help: Print message for pokedex
  map: Display names of location areas in Pokemon world
  mapb: Go back to previous location areas
  explore: List of all the pokemon located here. cmd:'explore location_name'
  catch: To catch your favourite Pokemon. cmd:'catch pokemon_name'
  inspect: Details of Pokemon. cmd:'inspect pokemon_name'
  pokedex: List of caught Pokemon
  exit: Exit the Pokedex
  ```

- **map**: Display names of location areas in the Pokémon world.  
  ```bash  
  ./pokedex map  
  ```

- **mapb**: Go back to the previous location areas.  
  ```bash  
  ./pokedex mapb  
  ```

- **explore**: List all Pokémon located in a specific area.  
  ```bash  
  ./pokedex explore <location_name>  
  ```

- **catch**: Catch a Pokémon and add it to your collection.  
  ```bash  
  ./pokedex catch <pokemon_name>  
  ```  

- **inspect**: View detailed information about a Pokémon.  
  ```bash  
  ./pokedex inspect <pokemon_name>  
  ```  

- **pokedex**: List all the Pokémon you have caught.  
  ```bash  
  ./pokedex pokedex  
  ```  

- **exit**: Exit the Pokédex application.  
  ```bash  
  ./pokedex exit  
  ```





## Contributing
Contributions are welcome! Feel free to fork the repository and submit a pull request.  

1. Fork the repository.  
2. Create a new branch:  
   ```bash  
   git checkout -b feature-branch-name  
   ```  
3. Commit your changes:  
   ```bash  
   git commit -m "Add feature"  
   ```  
4. Push to the branch:  
   ```bash  
   git push origin feature-branch-name  
   ```  
5. Open a pull request.  

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.  

## Acknowledgments
- [PokéAPI](https://pokeapi.co/) for providing Pokémon data.  
- Pokémon and all related content are © Nintendo, Game Freak, and The Pokémon Company.  
