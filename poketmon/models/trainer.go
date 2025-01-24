package models

import (
	"fmt"
)

const MaxEntryCount = 6

// Trainer represents a Pokemon trainer with a collection of Pokemon.
type Trainer struct {
	Name     string
	Pokemons []Pokemon
}

func (t *Trainer) CatchPokemon(pokemon Pokemon) {
	if len(t.Pokemons) < MaxEntryCount {
		pokemon.Trainer = t.Name
		t.Pokemons = append(t.Pokemons, pokemon)
	}
}

func (t *Trainer) PrintPokemons() {
	fmt.Printf("Trainer: %s\n", t.Name)
	fmt.Println("======================Print Pokemon======================")
	for _, pokemon := range t.Pokemons {
		fmt.Printf("Name: %-15s Level: %-3d Type: %-10s\n", pokemon.Name, pokemon.Level, pokemon.Type)

	}
	fmt.Println("========================================================")
}
