package pokemon

import "pokemon.com/models"

type SnorlaxFactory struct{}

func (s *SnorlaxFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "잠만보"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "몸통박치기", Damage: 50, PP: 35, Type: "노말"},
		{Name: "누르기", Damage: 80, PP: 15, Type: "노말"},
		{Name: "잠자기", Damage: 0, PP: 10, Type: "에스퍼"},
		{Name: "하이퍼빔", Damage: 150, PP: 5, Type: "노말"},
	}

	stats.Hp = 160
	stats.Attack = 110
	stats.Defense = 65
	stats.SpAtk = 65
	stats.SpDef = 110
	stats.Speed = 30

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
