package pokemon

import "pokemon.com/models"

type EkansFactory struct{}

func (e *EkansFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 22:
		name = "아보"
		evolution = "아보크"
		preEvolution = ""
		moves = []models.Move{
			{Name: "물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "독찌르기", Damage: 15, PP: 35, Type: "독"},
			{Name: "몸통박치기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "뱀파이어", Damage: 0, PP: 15, Type: "독"},
		}
		stats.Hp = 35
		stats.Attack = 60
		stats.Defense = 44
		stats.SpAtk = 40
		stats.SpDef = 54
		stats.Speed = 55
	default:
		name = "아보크"
		evolution = ""
		preEvolution = "아보"
		moves = []models.Move{
			{Name: "독가루", Damage: 0, PP: 20, Type: "독"},
			{Name: "독찌르기", Damage: 15, PP: 35, Type: "독"},
			{Name: "물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "뱀파이어", Damage: 0, PP: 15, Type: "독"},
		}
		stats.Hp = 60
		stats.Attack = 95
		stats.Defense = 69
		stats.SpAtk = 65
		stats.SpDef = 79
		stats.Speed = 80
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "독",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
