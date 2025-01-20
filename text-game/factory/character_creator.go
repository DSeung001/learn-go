package factory

import "text-game.com/character"

// CreateCharacter 함수
func CreateCharacter(name, class string, level int, factory ClassLevelFactory) *character.Character {
	c := &character.Character{
		Name:    name,
		Class:   class,
		Level:   level,
		HP:      100,
		Attack:  10, // 기본 공격력
		Defense: 5,  // 기본 방어력
	}
	c.Feature = factory.CreateFeature()
	c.Ability = factory.CreateAbility()
	c.Feature.ApplyEffect(c)
	return c
}
