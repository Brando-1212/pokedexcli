package main

import (
	"time"

	"github.com/Brando-1212/pokedexcli/internal/pokeapi"
	
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	repl(cfg)
}
