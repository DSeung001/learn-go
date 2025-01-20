package main

import (
	"fmt"
	"text-game.com/battle"
	"text-game.com/factory"
	"text-game.com/utils"
)

func main() {
	utils.InitRandomSeed()

	// 캐릭터 생성
	warriorFactory := factory.WarriorLevel1Factory{}
	wizardFactory := factory.WizardLevel1Factory{}

	warrior := factory.CreateCharacter("Arthur", "Warrior", 1, warriorFactory)
	wizard := factory.CreateCharacter("Merlin", "Wizard", 1, wizardFactory)

	// 캐릭터 출력
	fmt.Printf("%s has %s\n", warrior.Name, warrior.Feature.Description())
	fmt.Printf("%s can use %s\n", wizard.Name, wizard.Ability.Effect())

	// 전투
	battle.Battle(warrior, wizard)
}
