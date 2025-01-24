package pokemon

import "pokemon.com/models"

type GeodudeFactory struct {
	BaseStats map[string]int
}

func (g *GeodudeFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

	switch {
	case level < 25:
		name = "꼬마돌"
		evolution = "데구리"
		preEvolution = ""
		moves = []models.Move{
			{Name: "돌 던지기", Damage: 50, PP: 15, Type: "바위"},
			{Name: "머리받기", Damage: 70, PP: 15, Type: "노말"},
			{Name: "구르기", Damage: 30, PP: 20, Type: "바위"},
			{Name: "스톤샤워", Damage: 75, PP: 10, Type: "바위"},
		}
	case level < 40:
		name = "데구리"
		evolution = "딱구리"
		preEvolution = "꼬마돌"
		moves = []models.Move{
			{Name: "자이로볼", Damage: 80, PP: 5, Type: "강철"},
			{Name: "지진", Damage: 100, PP: 10, Type: "땅"},
			{Name: "스톤엣지", Damage: 100, PP: 5, Type: "바위"},
			{Name: "아이언 디펜스", Damage: 0, PP: 15, Type: "강철"},
		}
	default:
		name = "딱구리"
		evolution = ""
		preEvolution = "데구리"
		moves = []models.Move{
			{Name: "헤비슬램", Damage: 120, PP: 5, Type: "강철"},
			{Name: "폭발", Damage: 250, PP: 5, Type: "노말"},
			{Name: "파워젬", Damage: 80, PP: 20, Type: "바위"},
			{Name: "지구던지기", Damage: 50, PP: 20, Type: "땅"},
		}
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "바위/땅",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        g.BaseStats,
		Moves:        moves,
	}
}
