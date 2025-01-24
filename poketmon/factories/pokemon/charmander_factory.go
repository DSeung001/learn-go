package pokemon

import "pokemon.com/models"

type CharmanderFactory struct {
	BaseStats map[string]int
}

func (f *CharmanderFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

	switch {
	case level < 16:
		name = "파이리"
		evolution = "리자드"
		preEvolution = ""
		moves = []models.Move{
			{Name: "할퀴기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "불꽃세례", Damage: 40, PP: 25, Type: "불꽃"},
			{Name: "연막", Damage: 0, PP: 20, Type: "노말"},
			{Name: "용의숨결", Damage: 60, PP: 20, Type: "드래곤"},
		}
	case level < 36:
		name = "리자드"
		evolution = "리자몽"
		preEvolution = "파이리"
		moves = []models.Move{
			{Name: "화염엄니", Damage: 65, PP: 15, Type: "불꽃"},
			{Name: "불꽃펄스", Damage: 70, PP: 15, Type: "불꽃"},
			{Name: "베어가르기", Damage: 70, PP: 20, Type: "노말"},
			{Name: "드래곤크루", Damage: 80, PP: 15, Type: "드래곤"},
		}
	default:
		name = "리자몽"
		evolution = ""
		preEvolution = "리자드"
		moves = []models.Move{
			{Name: "화염방사", Damage: 90, PP: 15, Type: "불꽃"},
			{Name: "날기", Damage: 90, PP: 15, Type: "비행"},
			{Name: "열풍", Damage: 95, PP: 10, Type: "불꽃"},
			{Name: "인페르노", Damage: 100, PP: 5, Type: "불꽃"},
		}
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "불꽃",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        f.BaseStats,
		Moves:        moves,
	}
}
