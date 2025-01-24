package pokemon

import "pokemon.com/models"

type DittoFactory struct{}

func (d *DittoFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "메타몽"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "변신", Damage: 0, PP: 10, Type: "노말"},
	}
	stats.Hp = 48
	stats.Attack = 48
	stats.Defense = 48
	stats.SpAtk = 48
	stats.SpDef = 48
	stats.Speed = 48

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "노말",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
