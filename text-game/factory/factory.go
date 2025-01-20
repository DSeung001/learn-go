package factory

import "text-game.com/character"

// ClassLevelFactory 인터페이스 정의
type ClassLevelFactory interface {
	CreateFeature() character.Feature
	CreateAbility() character.Ability
}

// WarriorLevel1Factory 구현체
type WarriorLevel1Factory struct{}

func (w WarriorLevel1Factory) CreateFeature() character.Feature {
	return character.HeavyArmorProficiency{}
}

func (w WarriorLevel1Factory) CreateAbility() character.Ability {
	return character.SecondWind{}
}

// WizardLevel1Factory 구현체
type WizardLevel1Factory struct{}

func (w WizardLevel1Factory) CreateFeature() character.Feature {
	return character.Spellcasting{}
}

func (w WizardLevel1Factory) CreateAbility() character.Ability {
	return character.Fireball{}
}
