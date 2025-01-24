package pokemon

import "pokemon.com/models"

type MewtwoFactory struct{}

func (m *MewtwoFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "뮤츠"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "사이코 컷", Damage: 70, PP: 20, Type: "에스퍼"},
		{Name: "사이코 킬", Damage: 100, PP: 10, Type: "에스퍼"},
		{Name: "하이퍼빔", Damage: 150, PP: 5, Type: "노말"},
		{Name: "명상", Damage: 0, PP: 20, Type: "에스퍼"},
	}

	stats.Hp = 106
	stats.Attack = 110
	stats.Defense = 90
	stats.SpAtk = 154
	stats.SpDef = 90
	stats.Speed = 130

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
