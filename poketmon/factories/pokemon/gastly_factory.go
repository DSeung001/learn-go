package pokemon

import "pokemon.com/models"

type GastlyFactory struct{}

func (g *GastlyFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 25:
		name = "고오스"
		evolution = "고우스트"
		preEvolution = ""
		moves = []models.Move{
			{Name: "악몽", Damage: 0, PP: 10, Type: "고스트"},
			{Name: "환영빔", Damage: 60, PP: 15, Type: "고스트"},
			{Name: "최면술", Damage: 0, PP: 20, Type: "에스퍼"},
			{Name: "독가루", Damage: 0, PP: 35, Type: "독"},
		}
		stats.Hp = 30
		stats.Attack = 35
		stats.Defense = 30
		stats.SpAtk = 100
		stats.SpDef = 35
		stats.Speed = 80
	case level < 40:
		name = "고우스트"
		evolution = "팬텀"
		preEvolution = "고오스"
		moves = []models.Move{
			{Name: "악몽", Damage: 0, PP: 10, Type: "고스트"},
			{Name: "환영빔", Damage: 75, PP: 10, Type: "고스트"},
			{Name: "그림자 구슬", Damage: 80, PP: 15, Type: "고스트"},
			{Name: "사념파", Damage: 90, PP: 10, Type: "에스퍼"},
		}
		stats.Hp = 45
		stats.Attack = 50
		stats.Defense = 45
		stats.SpAtk = 115
		stats.SpDef = 55
		stats.Speed = 95
	default:
		name = "팬텀"
		evolution = ""
		preEvolution = "고우스트"
		moves = []models.Move{
			{Name: "섀도우볼", Damage: 100, PP: 10, Type: "고스트"},
			{Name: "드림이터", Damage: 100, PP: 15, Type: "에스퍼"},
			{Name: "섀도우펀치", Damage: 90, PP: 15, Type: "고스트"},
			{Name: "고스트 다이브", Damage: 120, PP: 5, Type: "고스트"},
		}
		stats.Hp = 60
		stats.Attack = 65
		stats.Defense = 60
		stats.SpAtk = 130
		stats.SpDef = 75
		stats.Speed = 110
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "고스트",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
