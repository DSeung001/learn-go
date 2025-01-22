package factory

import "text-game.com/character"

// ClassLevelFactory 인터페이스 정의
type ClassLevelFactory interface {
	CreateFeature() character.Feature
	CreateAbility() character.Ability
	CreateBaseStats() (hp int, attack int, defense int)
}

// WarriorLevel1Factory 구현체
type WarriorLevel1Factory struct{}

func (w WarriorLevel1Factory) CreateFeature() character.Feature {
	return character.HeavyArmorProficiency{}
}

func (w WarriorLevel1Factory) CreateAbility() character.Ability {
	return character.SecondWind{}
}

func (w WarriorLevel1Factory) CreateBaseStats() (hp int, attack int, defense int) {
	return 50, 12, 9
}

// WizardLevel1Factory 구현체
type WizardLevel1Factory struct{}

func (w WizardLevel1Factory) CreateFeature() character.Feature {
	return character.Spellcasting{}
}

func (w WizardLevel1Factory) CreateAbility() character.Ability {
	return character.Fireball{}
}

func (w WizardLevel1Factory) CreateBaseStats() (hp int, attack int, defense int) {
	return 35, 17, 5
}
