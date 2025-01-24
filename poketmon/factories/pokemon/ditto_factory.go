package pokemon

import "pokemon.com/models"

type DittoFactory struct {
	BaseStats map[string]int
}

func (d *DittoFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

	name = "메타몽"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "변신", Damage: 0, PP: 10, Type: "노말"},
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "노말",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        d.BaseStats,
		Moves:        moves,
	}
}
