package pokemon

import "pokemon.com/models"

type RayquazaFactory struct{}

func (r *RayquazaFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "레쿠쟈"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "용의춤", Damage: 0, PP: 20, Type: "용"},
		{Name: "공중날기", Damage: 90, PP: 15, Type: "비행"},
		{Name: "용의분노", Damage: 40, PP: 10, Type: "용"},
		{Name: "폭풍", Damage: 120, PP: 5, Type: "비행"},
	}
	stats.Hp = 105
	stats.Attack = 150
	stats.Defense = 90
	stats.SpAtk = 150
	stats.SpDef = 90
	stats.Speed = 95

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "용/비행",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
