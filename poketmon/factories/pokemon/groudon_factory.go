package pokemon

import "pokemon.com/models"

type GroudonFactory struct{}

func (g *GroudonFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "그란돈"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "지진", Damage: 100, PP: 10, Type: "땅"},
		{Name: "솔라빔", Damage: 120, PP: 10, Type: "풀"},
		{Name: "불대문자", Damage: 110, PP: 5, Type: "불꽃"},
		{Name: "스톤엣지", Damage: 100, PP: 5, Type: "바위"},
	}
	stats.Hp = 100
	stats.Attack = 150
	stats.Defense = 140
	stats.SpAtk = 100
	stats.SpDef = 90
	stats.Speed = 90

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "땅",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
