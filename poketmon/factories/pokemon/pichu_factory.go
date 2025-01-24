package pokemon

import "pokemon.com/models"

type PichuFactory struct{}

func (e *PichuFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 15:
		name = "피츄"
		evolution = "피카츄"
		preEvolution = ""
		moves = []models.Move{
			{Name: "전기 쇼크", Damage: 40, PP: 30, Type: "전기"},
			{Name: "매력", Damage: 0, PP: 20, Type: "페어리"},
			{Name: "달콤한 키스", Damage: 0, PP: 10, Type: "페어리"},
			{Name: "볼트태클", Damage: 20, PP: 20, Type: "전기"},
		}
		stats.Hp = 20
		stats.Attack = 40
		stats.Defense = 15
		stats.SpAtk = 35
		stats.SpDef = 35
		stats.Speed = 60
	case level < 30:
		name = "피카츄"
		evolution = "라이츄"
		preEvolution = "피츄"
		moves = []models.Move{
			{Name: "10만 볼트", Damage: 90, PP: 15, Type: "전기"},
			{Name: "전광석화", Damage: 40, PP: 30, Type: "노말"},
			{Name: "전기 볼", Damage: 60, PP: 10, Type: "전기"},
			{Name: "속임수", Damage: 30, PP: 10, Type: "노말"},
		}
		stats.Hp = 35
		stats.Attack = 55
		stats.Defense = 40
		stats.SpAtk = 50
		stats.SpDef = 50
		stats.Speed = 90
	default:
		name = "라이츄"
		evolution = ""
		preEvolution = "피카츄"
		moves = []models.Move{
			{Name: "번개 펀치", Damage: 75, PP: 15, Type: "전기"},
			{Name: "볼트태클", Damage: 120, PP: 5, Type: "전기"},
			{Name: "아이언 테일", Damage: 100, PP: 15, Type: "강철"},
			{Name: "방전", Damage: 80, PP: 15, Type: "전기"},
		}
		stats.Hp = 60
		stats.Attack = 90
		stats.Defense = 55
		stats.SpAtk = 90
		stats.SpDef = 80
		stats.Speed = 110
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "전기",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
