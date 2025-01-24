package battles

import (
	"fmt"
	"math/rand"
	"pokemon.com/models"
	"time"
)

// BattlePokemon handles a battle between two Pokemon
func BattlePokemon(pokemon1, pokemon2 *models.Pokemon) string {
	fmt.Printf("\nBattle Start! %s vs %s\n", pokemon1.Name+"["+pokemon1.Trainer+"]", pokemon2.Name+"["+pokemon2.Trainer+"]")

	for pokemon1.Stats.Hp > 0 && pokemon2.Stats.Hp > 0 {
		// Determine turn order based on Speed
		var first, second *models.Pokemon
		if pokemon1.Stats.Speed >= pokemon2.Stats.Speed {
			first, second = pokemon1, pokemon2
		} else {
			first, second = pokemon2, pokemon1
		}

		// First Pokemon attacks
		damage := calculateDamage(first, second)
		second.Stats.Hp -= damage
		fmt.Printf("%s attacks %s for %d damage! (%s HP: %d)\n",
			first.Name, second.Name, damage, second.Name, second.Stats.Hp)

		if second.Stats.Hp <= 0 {
			fmt.Printf("%s fainted!\n", second.Name)
			return first.Name
		}

		// Second Pokemon attacks
		damage = calculateDamage(second, first)
		first.Stats.Hp -= damage
		fmt.Printf("%s attacks %s for %d damage! (%s HP: %d)\n",
			second.Name, first.Name, damage, first.Name, first.Stats.Hp)

		if first.Stats.Hp <= 0 {
			fmt.Printf("%s fainted!\n", first.Name)
			return second.Name
		}
	}

	return ""
}

// 나중에 live 랑 current Hp 로직에 적용하기
func BattleTrainers(trainer1, trainer2 *models.Trainer) string {
	fmt.Printf("\nBattle Start! %s vs %s\n", trainer1.Name, trainer2.Name)

	// Each trainer uses their Pokemon in order
	for len(trainer1.Pokemons) > 0 && len(trainer2.Pokemons) > 0 {
		// Get the first Pokemon from each trainer
		pokemon1 := &trainer1.Pokemons[0]
		pokemon2 := &trainer2.Pokemons[0]

		// Battle between the two Pokemon
		winner := BattlePokemon(pokemon1, pokemon2)

		// Remove the fainted Pokemon
		if winner == pokemon1.Name {
			trainer2.Pokemons = trainer2.Pokemons[1:]
		} else {
			trainer1.Pokemons = trainer1.Pokemons[1:]
		}
	}

	// Determine the winner
	if len(trainer1.Pokemons) > 0 {
		fmt.Printf("\n%s wins the battle! with %d Pokemon remaining!\n", trainer1.Name, len(trainer1.Pokemons))
		return trainer1.Name
	}

	fmt.Printf("\n%s wins the battle! with %d Pokemon remaining!\n", trainer2.Name, len(trainer2.Pokemons))
	return trainer2.Name
}

// Damage 계산 함수
func calculateDamage(attacker, defensor *models.Pokemon) int {
	rand.Seed(time.Now().UnixNano())

	chosenMove := attacker.Moves[rand.Intn(len(attacker.Moves))]
	power := float64(chosenMove.Damage)
	if power == 0 {
		power = 20
	}
	fmt.Printf("%s, %s\n", attacker.Name+"["+attacker.Trainer+"]", chosenMove.Name)

	isSpecialAttack := attacker.Stats.SpAtk > attacker.Stats.Attack
	var attackFactor float64
	var defenseFactor float64

	if isSpecialAttack {
		attackFactor = float64(attacker.Stats.SpAtk)  // Special Attack 사용
		defenseFactor = float64(defensor.Stats.SpDef) // Special Defense 사용
	} else {
		attackFactor = float64(attacker.Stats.Attack)   // 일반 Attack 사용
		defenseFactor = float64(defensor.Stats.Defense) // 일반 Defense 사용
	}

	levelFactor := (2*float64(attacker.Level))/5.0 + 2
	modifier := 1.0

	// Damage 계산 공식
	baseDamage := (((levelFactor * power * attackFactor / defenseFactor) / 50) + 2) * modifier
	return int(baseDamage)
}
