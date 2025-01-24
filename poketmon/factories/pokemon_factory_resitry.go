package factories

import (
	"errors"
	"math/rand"
	"pokemon.com/factories/pokemon"
)

type PokemonFactoryRegistry struct {
	pokemonFactories []PokemonFactory
}

func NewPokemonFactoryRegistry() *PokemonFactoryRegistry {
	return &PokemonFactoryRegistry{
		pokemonFactories: []PokemonFactory{},
	}
}

func (r *PokemonFactoryRegistry) Register(factory PokemonFactory) {
	r.pokemonFactories = append(r.pokemonFactories, factory)
}

func (r *PokemonFactoryRegistry) GetRandomPokemonFactory() (PokemonFactory, error) {
	if len(r.pokemonFactories) == 0 {
		return nil, errors.New("no factories registered")
	}
	randomIndex := rand.Intn(len(r.pokemonFactories))
	return r.pokemonFactories[randomIndex], nil
}

func (r *PokemonFactoryRegistry) PokemonFactoryRegister() {
	r.Register(&pokemon.BulbasaurFactory{})
	r.Register(&pokemon.CharmanderFactory{})
	r.Register(&pokemon.DittoFactory{})
	r.Register(&pokemon.GastlyFactory{})
	r.Register(&pokemon.GeodudeFactory{})
	r.Register(&pokemon.LugiaFactory{})
	r.Register(&pokemon.MagnemiteFactory{})
	r.Register(&pokemon.MewtwoFactory{})
	r.Register(&pokemon.PichuFactory{})
	r.Register(&pokemon.SnorlaxFactory{})
	r.Register(&pokemon.SquirtleFactory{})
	r.Register(&pokemon.VulpixFactory{})
	r.Register(&pokemon.ZubatFactory{})
}
