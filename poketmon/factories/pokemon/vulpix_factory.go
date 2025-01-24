package pokemon

import "pokemon.com/models"

type VulpixFactory struct {
	BaseStats map[string]int
}

func (v *VulpixFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

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
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "불꽃",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        v.BaseStats,
		Moves:        moves,
	}
}
