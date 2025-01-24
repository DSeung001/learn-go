package pokemon

import "pokemon.com/models"

type KyogreFactory struct{}

func (k *KyogreFactory) NewPokemon(level int) models.Pokemon {
	var name, evolution, preEvolution string
	var moves []models.Move
	var stats = models.Stats{}

	name = "가이오가"
	evolution = ""
	preEvolution = ""
	moves = []models.Move{
		{Name: "하이드로펌프", Damage: 110, PP: 5, Type: "물"},
		{Name: "냉동빔", Damage: 90, PP: 10, Type: "얼음"},
		{Name: "파도타기", Damage: 90, PP: 15, Type: "물"},
		{Name: "폭풍우", Damage: 120, PP: 5, Type: "물"},
	}
	stats.Hp = 100
	stats.Attack = 100
	stats.Defense = 90
	stats.SpAtk = 150
	stats.SpDef = 140
	stats.Speed = 90

	return models.Pokemon{
		Level:        level,
		Name:         name,
		Type:         "물",
		Evolution:    evolution,
		PreEvolution: preEvolution,
		Stats:        stats,
		Moves:        moves,
	}
}
