package battle

import (
	"fmt"
	"text-game.com/character"
)

// Battle 함수
func Battle(c1, c2 *character.Character) {
	fmt.Printf("Battle Start: %s vs %s!\n", c1.Name, c2.Name)
	round := 1
	for !c1.IsDefeated && !c2.IsDefeated {
		fmt.Printf("\n--- Round %d ---\n", round)
		Attack(c1, c2)
		if c2.IsDefeated {
			fmt.Printf("%s wins!\n", c1.Name)
			break
		}
		Attack(c2, c1)
		if c1.IsDefeated {
			fmt.Printf("%s wins!\n", c2.Name)
			break
		}
		round++
	}
}

// Attack 함수
func Attack(attacker, defender *character.Character) {
	damage := attacker.Attack - defender.Defense
	if damage < 0 {
		damage = 0
	}
	defender.HP -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", attacker.Name, defender.Name, damage)
	if defender.HP <= 0 {
		defender.IsDefeated = true
		fmt.Printf("%s is defeated!\n", defender.Name)
	}
}
