package factories

import (
	"math/rand"
	"pokemon.com/models"
)

// TrainerFactory defines the interface for creating trainers.
type TrainerFactory interface {
	SetTrainerName(name string, pokemonFactories []PokemonFactory) models.Trainer
}

// DefaultTrainerFactory implements TrainerFactory.
type DefaultTrainerFactory struct{}

func (t *DefaultTrainerFactory) SetTrainerName(name string, pokemonFactories []PokemonFactory) models.Trainer {
	var pokemons []models.Pokemon
	for _, factory := range pokemonFactories {
		pokemons = append(pokemons, factory.NewPokemon(rand.Intn(50)+1))
	}
	return models.Trainer{
		Name:     name,
		Pokemons: pokemons,
	}
}
