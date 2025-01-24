package pokemon

import "pokemon.com/models"

type StarlyFactory struct{}

func (s *StarlyFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	switch {
	case level < 14:
		name = "찌르꼬"
		evolution = "찌르버드"
		preEvolution = ""
		moves = []models.Move{
			{Name: "날개치기", Damage: 35, PP: 35, Type: "비행"},
			{Name: "울음소리", Damage: 0, PP: 40, Type: "노말"},
			{Name: "몸통박치기", Damage: 40, PP: 35, Type: "노말"},
			{Name: "쪼기", Damage: 35, PP: 35, Type: "비행"},
		}
		stats.Hp = 40
		stats.Attack = 55
		stats.Defense = 30
		stats.SpAtk = 30
		stats.SpDef = 30
		stats.Speed = 60
	default:
		name = "찌르버드"
		evolution = "찌르호크"
		preEvolution = "찌르꼬"
		moves = []models.Move{
			{Name: "공중날기", Damage: 90, PP: 15, Type: "비행"},
			{Name: "쪼기", Damage: 35, PP: 35, Type: "비행"},
			{Name: "제비반환", Damage: 60, PP: 20, Type: "비행"},
			{Name: "날개치기", Damage: 35, PP: 35, Type: "비행"},
		}
		stats.Hp = 55
		stats.Attack = 75
		stats.Defense = 50
		stats.SpAtk = 40
		stats.SpDef = 40
		stats.Speed = 80
	}

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "노말/비행",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
