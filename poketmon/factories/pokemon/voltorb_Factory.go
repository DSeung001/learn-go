package pokemon

import "pokemon.com/models"

type VoltorbFactory struct{}

func (v *VoltorbFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 30:
		name = "찌리리공"
		evolution = "붐볼"
		preEvolution = ""
		moves = []models.Move{
			{Name: "전기 쇼크", Damage: 40, PP: 30, Type: "전기"},
			{Name: "스파크", Damage: 65, PP: 20, Type: "전기"},
			{Name: "라이트스크린", Damage: 0, PP: 20, Type: "에스퍼"},
			{Name: "자폭", Damage: 200, PP: 5, Type: "노말"},
		}
		stats.Hp = 40
		stats.Attack = 30
		stats.Defense = 50
		stats.SpAtk = 55
		stats.SpDef = 55
		stats.Speed = 100
	default:
		name = "붐볼"
		evolution = ""
		preEvolution = "찌리리공"
		moves = []models.Move{
			{Name: "썬더볼트", Damage: 90, PP: 15, Type: "전기"},
			{Name: "전자포", Damage: 120, PP: 5, Type: "전기"},
			{Name: "스위프트", Damage: 60, PP: 20, Type: "노말"},
			{Name: "자폭", Damage: 200, PP: 5, Type: "노말"},
		}
		stats.Hp = 60
		stats.Attack = 50
		stats.Defense = 70
		stats.SpAtk = 80
		stats.SpDef = 80
		stats.Speed = 150
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
