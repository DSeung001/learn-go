package pokemon

import "pokemon.com/models"

type ZubatFactory struct {
	BaseStats map[string]int
}

func (z *ZubatFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move

	switch {
	case level < 22:
		name = "주뱃"
		evolution = "골뱃"
		preEvolution = ""
		moves = []models.Move{
			{Name: "물기", Damage: 60, PP: 25, Type: "악"},
			{Name: "초음파", Damage: 0, PP: 20, Type: "노말"},
			{Name: "날개치기", Damage: 60, PP: 35, Type: "비행"},
			{Name: "독침", Damage: 15, PP: 35, Type: "독"},
		}
	case level < 40:
		name = "골뱃"
		evolution = "크로뱃"
		preEvolution = "주뱃"
		moves = []models.Move{
			{Name: "에어슬래시", Damage: 75, PP: 15, Type: "비행"},
			{Name: "독가루", Damage: 0, PP: 35, Type: "독"},
			{Name: "스틸윙", Damage: 70, PP: 25, Type: "강철"},
			{Name: "섀도우볼", Damage: 80, PP: 15, Type: "고스트"},
		}
	default:
		name = "크로뱃"
		evolution = ""
		preEvolution = "골뱃"
		moves = []models.Move{
			{Name: "브레이브버드", Damage: 120, PP: 15, Type: "비행"},
			{Name: "크로스포이즌", Damage: 70, PP: 20, Type: "독"},
			{Name: "하이퍼빔", Damage: 150, PP: 5, Type: "노말"},
			{Name: "에어컷터", Damage: 60, PP: 25, Type: "비행"},
		}
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "비행/독",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        z.BaseStats,
		Moves:        moves,
	}
}
