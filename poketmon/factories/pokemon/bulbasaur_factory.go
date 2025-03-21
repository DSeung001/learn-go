package pokemon

import "pokemon.com/models"

type BulbasaurFactory struct{}

func (g *BulbasaurFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 16:
		name = "이상해씨"
		evolution = "이상해풀"
		preEvolution = ""
		moves = []models.Move{
			{Name: "몸통박치기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "덩굴채찍", Damage: 45, PP: 25, Type: "풀"},
			{Name: "씨뿌리기", Damage: 0, PP: 10, Type: "풀"},
			{Name: "독가루", Damage: 0, PP: 35, Type: "독"},
		}
		stats.Hp = 45
		stats.Attack = 49
		stats.Defense = 49
		stats.SpAtk = 65
		stats.SpDef = 65
		stats.Speed = 45

	case level < 32:
		name = "이상해풀"
		evolution = "이상해꽃"
		preEvolution = "이상해씨"
		moves = []models.Move{
			{Name: "잎날가르기", Damage: 55, PP: 25, Type: "풀"},
			{Name: "수면가루", Damage: 0, PP: 15, Type: "풀"},
			{Name: "씨폭탄", Damage: 80, PP: 15, Type: "풀"},
			{Name: "몸통박치기", Damage: 90, PP: 20, Type: "노말"},
		}
		stats.Hp = 60
		stats.Attack = 62
		stats.Defense = 63
		stats.SpAtk = 80
		stats.SpDef = 80
		stats.Speed = 60
	default:
		name = "이상해꽃" // Venusaur
		evolution = ""
		preEvolution = "이상해풀" // Ivysaur
		moves = []models.Move{
			{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
			{Name: "오물폭탄", Damage: 90, PP: 10, Type: "독"},
			{Name: "지진", Damage: 100, PP: 10, Type: "땅"},
			{Name: "기가드레인", Damage: 75, PP: 10, Type: "풀"},
		}
		stats.Hp = 80
		stats.Attack = 82
		stats.Defense = 83
		stats.SpAtk = 100
		stats.SpDef = 100
		stats.Speed = 80
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "풀",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
