package main

import (
	"fmt"
	
)
func commandPokedex(cfg *config, arg ...string) error{

	fmt.Println("Your Pokedex")
	for _, pok := range cfg.pokedex {
		fmt.Printf(" - %s\n", pok.Name)
	}
	return nil
}