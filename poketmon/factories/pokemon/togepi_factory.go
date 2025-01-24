package pokemon

import "pokemon.com/models"

type TogepiFactory struct{}

func (t *TogepiFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 20:
		name = "토개피"
		evolution = "토게틱"
		preEvolution = ""
		moves = []models.Move{
			{Name: "마법의잎사귀", Damage: 60, PP: 20, Type: "풀"},
			{Name: "매직코트", Damage: 0, PP: 15, Type: "에스퍼"},
			{Name: "하품", Damage: 0, PP: 10, Type: "노말"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
		}
		stats.Hp = 35
		stats.Attack = 20
		stats.Defense = 65
		stats.SpAtk = 40
		stats.SpDef = 65
		stats.Speed = 20
	case level < 40:
		name = "토게틱"
		evolution = "토게키스"
		preEvolution = "토개피"
		moves = []models.Move{
			{Name: "공중날기", Damage: 90, PP: 15, Type: "비행"},
			{Name: "매직샤인", Damage: 80, PP: 10, Type: "페어리"},
			{Name: "치유소원", Damage: 0, PP: 10, Type: "에스퍼"},
			{Name: "사이코키네시스", Damage: 90, PP: 10, Type: "에스퍼"},
		}
		stats.Hp = 55
		stats.Attack = 40
		stats.Defense = 85
		stats.SpAtk = 80
		stats.SpDef = 105
		stats.Speed = 40
	default:
		name = "토게키스"
		evolution = ""
		preEvolution = "토게틱"
		moves = []models.Move{
			{Name: "매지컬샤인", Damage: 80, PP: 10, Type: "페어리"},
			{Name: "하이퍼빔", Damage: 150, PP: 5, Type: "노말"},
			{Name: "에어슬래쉬", Damage: 75, PP: 15, Type: "비행"},
			{Name: "파동탄", Damage: 80, PP: 20, Type: "격투"},
		}
		stats.Hp = 85
		stats.Attack = 50
		stats.Defense = 95
		stats.SpAtk = 120
		stats.SpDef = 115
		stats.Speed = 80
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "페어리/비행",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
