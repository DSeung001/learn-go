package pokemon

import "pokemon.com/models"

type LaprasFactory struct{}

func (l *LaprasFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "라프라스"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "얼다바람", Damage: 60, PP: 15, Type: "얼음"},
		{Name: "하이드로펌프", Damage: 110, PP: 5, Type: "물"},
		{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
		{Name: "파동탄", Damage: 80, PP: 20, Type: "격투"},
	}
	stats.Hp = 130
	stats.Attack = 85
	stats.Defense = 80
	stats.SpAtk = 85
	stats.SpDef = 95
	stats.Speed = 60

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "물/얼음",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
