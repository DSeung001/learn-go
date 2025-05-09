package pokemon

import "pokemon.com/models"

type MagnemiteFactory struct{}

func (m *MagnemiteFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 30:
		name = "코일"
		evolution = "레어코일"
		preEvolution = ""
		moves = []models.Move{
			{Name: "전기 쇼크", Damage: 40, PP: 30, Type: "전기"},
			{Name: "스파크", Damage: 65, PP: 20, Type: "전기"},
			{Name: "금속 음파", Damage: 0, PP: 40, Type: "강철"},
			{Name: "방전", Damage: 80, PP: 15, Type: "전기"},
		}
		stats.Hp = 25
		stats.Attack = 35
		stats.Defense = 70
		stats.SpAtk = 95
		stats.SpDef = 55
		stats.Speed = 45
	case level < 50:
		name = "레어코일"
		evolution = "자포코일"
		preEvolution = "코일"
		moves = []models.Move{
			{Name: "전자포", Damage: 120, PP: 5, Type: "전기"},
			{Name: "자속포", Damage: 90, PP: 10, Type: "전기"},
			{Name: "미러 샷", Damage: 65, PP: 20, Type: "강철"},
			{Name: "라이트스크린", Damage: 0, PP: 20, Type: "에스퍼"},
		}
		stats.Hp = 50
		stats.Attack = 60
		stats.Defense = 95
		stats.SpAtk = 120
		stats.SpDef = 70
		stats.Speed = 70
	default:
		name = "자포코일"
		evolution = ""
		preEvolution = "레어코일"
		moves = []models.Move{
			{Name: "썬더", Damage: 110, PP: 10, Type: "전기"},
			{Name: "트라이어택", Damage: 80, PP: 10, Type: "노말"},
			{Name: "플래시 캐논", Damage: 90, PP: 10, Type: "강철"},
			{Name: "매그넷 봄", Damage: 60, PP: 20, Type: "강철"},
		}
		stats.Hp = 70
		stats.Attack = 70
		stats.Defense = 115
		stats.SpAtk = 130
		stats.SpDef = 90
		stats.Speed = 60
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "전기/강철",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
