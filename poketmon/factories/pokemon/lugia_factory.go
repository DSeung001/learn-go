package pokemon

import "pokemon.com/models"

type LugiaFactory struct{}

func (l *LugiaFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "루기아"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "에어로블래스트", Damage: 120, PP: 5, Type: "비행"},
		{Name: "명상", Damage: 0, PP: 20, Type: "에스퍼"},
		{Name: "하이드로펌프", Damage: 110, PP: 5, Type: "물"},
		{Name: "리플렉터", Damage: 0, PP: 20, Type: "에스퍼"},
	}

	stats.Hp = 106
	stats.Attack = 90
	stats.Defense = 130
	stats.SpAtk = 90
	stats.SpDef = 154
	stats.Speed = 110

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "에스퍼",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
