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
	if len(t.Pokemons) <= MaxEntryCount {
		t.Pokemons = append(t.Pokemons, pokemon)
	}
}

func (t *Trainer) PrintPokemons() {
	fmt.Printf("Trainer: %s\n", t.Name)
	fmt.Println("======================Print Pokemon======================")
	for _, pokemon := range t.Pokemons {
		fmt.Printf("Name: %-15s Level: %-3d Type: %-10s\n", pokemon.Name, pokemon.Level, pokemon.Type)
		fmt.Printf("  Evolution: %-10s Pre-Evolution: %-10s\n", pokemon.Evolution, pokemon.PreEvolution)
		fmt.Println("  Stats:")
		for stat, value := range pokemon.Stats {
			fmt.Printf("    %-10s: %d\n", stat, value)
		}
		fmt.Println("  Moves:")
		for _, move := range pokemon.Moves {
			fmt.Printf("    %-15s (Type: %-10s, Damage: %d, PP: %d)\n", move.Name, move.Type, move.Damage, move.PP)
		}
		fmt.Println("--------------------------------------------------------")
	}
	fmt.Println("========================================================")
}
