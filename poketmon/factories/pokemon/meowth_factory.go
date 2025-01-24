package pokemon

import "pokemon.com/models"

type MeowthFactory struct{}

func (m *MeowthFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 28:
		name = "냐옹"
		evolution = "페르시온"
		preEvolution = ""
		moves = []models.Move{
			{Name: "물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "할퀴기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "제비반환", Damage: 60, PP: 20, Type: "비행"},
			{Name: "도둑질", Damage: 60, PP: 25, Type: "악"},
		}
		stats.Hp = 40
		stats.Attack = 45
		stats.Defense = 35
		stats.SpAtk = 40
		stats.SpDef = 40
		stats.Speed = 90
	default:
		name = "페르시온"
		evolution = ""
		preEvolution = "냐옹"
		moves = []models.Move{
			{Name: "물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "베어가르기", Damage: 70, PP: 20, Type: "노말"},
			{Name: "할퀴기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "고속이동", Damage: 0, PP: 20, Type: "에스퍼"},
		}
		stats.Hp = 65
		stats.Attack = 70
		stats.Defense = 60
		stats.SpAtk = 65
		stats.SpDef = 65
		stats.Speed = 115
	}

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
