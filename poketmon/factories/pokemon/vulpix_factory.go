package pokemon

import "pokemon.com/models"

type VulpixFactory struct{}

func (v *VulpixFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 25:
		name = "식스테일"
		evolution = "나인테일"
		preEvolution = ""
		moves = []models.Move{
			{Name: "불꽃세례", Damage: 40, PP: 25, Type: "불꽃"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
			{Name: "용해액", Damage: 40, PP: 30, Type: "독"},
			{Name: "윌오위스프", Damage: 0, PP: 15, Type: "불꽃"},
		}
		stats.Hp = 38
		stats.Attack = 41
		stats.Defense = 40
		stats.SpAtk = 50
		stats.SpDef = 65
		stats.Speed = 65
	default:
		name = "나인테일"
		evolution = ""
		preEvolution = "식스테일"
		moves = []models.Move{
			{Name: "불대문자", Damage: 110, PP: 5, Type: "불꽃"},
			{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
			{Name: "냉동빔", Damage: 90, PP: 10, Type: "얼음"},
			{Name: "명상", Damage: 0, PP: 20, Type: "에스퍼"},
		}
		stats.Hp = 73
		stats.Attack = 76
		stats.Defense = 75
		stats.SpAtk = 81
		stats.SpDef = 100
		stats.Speed = 100
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
