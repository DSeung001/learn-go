package pokemon

import "pokemon.com/models"

type GastlyFactory struct {
	BaseStats map[string]int
}

func (g *GastlyFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

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
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "고스트",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        g.BaseStats,
		Moves:        moves,
	}
}
