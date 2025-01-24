package pokemon

import "pokemon.com/models"

type PonytaFactory struct{}

func (p *PonytaFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 40:
		name = "포니타"
		evolution = "날쌩마"
		preEvolution = ""
		moves = []models.Move{
			{Name: "불꽃세례", Damage: 40, PP: 25, Type: "불꽃"},
			{Name: "몸통박치기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
			{Name: "스모그", Damage: 30, PP: 20, Type: "독"},
		}
		stats.Hp = 50
		stats.Attack = 85
		stats.Defense = 55
		stats.SpAtk = 65
		stats.SpDef = 65
		stats.Speed = 90
	default:
		name = "날쌩마"
		evolution = ""
		preEvolution = "포니타"
		moves = []models.Move{
			{Name: "불대문자", Damage: 110, PP: 5, Type: "불꽃"},
			{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
			{Name: "사이코키네시스", Damage: 90, PP: 10, Type: "에스퍼"},
			{Name: "화염방사", Damage: 90, PP: 15, Type: "불꽃"},
		}
		stats.Hp = 65
		stats.Attack = 100
		stats.Defense = 70
		stats.SpAtk = 80
		stats.SpDef = 80
		stats.Speed = 105
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "불꽃",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
