package pokemon

import "pokemon.com/models"

type SquirtleFactory struct{}

func (w *SquirtleFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 16:
		name = "꼬부기"
		evolution = "어니부기"
		preEvolution = ""
		moves = []models.Move{
			{Name: "몸통박치기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "물대포", Damage: 40, PP: 25, Type: "물"},
			{Name: "껍질에 숨기", Damage: 0, PP: 40, Type: "물"},
			{Name: "거품", Damage: 40, PP: 30, Type: "물"},
		}
		stats.Hp = 44
		stats.Attack = 48
		stats.Defense = 65
		stats.SpAtk = 50
		stats.SpDef = 64
		stats.Speed = 43
	case level < 36:
		name = "어니부기"
		evolution = "거북왕"
		preEvolution = "꼬부기"
		moves = []models.Move{
			{Name: "깨물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "파도타기", Damage: 60, PP: 20, Type: "물"},
			{Name: "아쿠아테일", Damage: 90, PP: 10, Type: "물"},
			{Name: "방어", Damage: 0, PP: 10, Type: "노말"},
		}
		stats.Hp = 59
		stats.Attack = 63
		stats.Defense = 80
		stats.SpAtk = 65
		stats.SpDef = 80
		stats.Speed = 58
	default:
		name = "거북왕"
		evolution = ""
		preEvolution = "어니부기"
		moves = []models.Move{
			{Name: "하이드로펌프", Damage: 110, PP: 5, Type: "물"},
			{Name: "로켓헤드", Damage: 130, PP: 10, Type: "노말"},
			{Name: "냉동빔", Damage: 90, PP: 10, Type: "얼음"},
			{Name: "서핑", Damage: 90, PP: 15, Type: "물"},
		}
		stats.Hp = 79
		stats.Attack = 83
		stats.Defense = 100
		stats.SpAtk = 85
		stats.SpDef = 105
		stats.Speed = 78
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "물",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
