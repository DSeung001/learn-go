package factories

import "pokemon.com/models"

type PokemonFactory interface {
	NewPokemon(level int) models.Pokemon
}
