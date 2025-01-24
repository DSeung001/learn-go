package pokemon

import "pokemon.com/models"

type ExeggcuteFactory struct{}

func (e *ExeggcuteFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 25:
		name = "아라리"
		evolution = "나시"
		preEvolution = ""
		moves = []models.Move{
			{Name: "사이코키네시스", Damage: 90, PP: 10, Type: "에스퍼"},
			{Name: "씨폭탄", Damage: 80, PP: 15, Type: "풀"},
			{Name: "하품", Damage: 0, PP: 10, Type: "노말"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
		}
		stats.Hp = 60
		stats.Attack = 40
		stats.Defense = 80
		stats.SpAtk = 60
		stats.SpDef = 45
		stats.Speed = 40
	default:
		name = "나시"
		evolution = ""
		preEvolution = "아라리"
		moves = []models.Move{
			{Name: "사이코키네시스", Damage: 90, PP: 10, Type: "에스퍼"},
			{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
			{Name: "씨폭탄", Damage: 80, PP: 15, Type: "풀"},
		}
		stats.Hp = 95
		stats.Attack = 95
		stats.Defense = 85
		stats.SpAtk = 125
		stats.SpDef = 75
		stats.Speed = 55
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "풀/에스퍼",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
