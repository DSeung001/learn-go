package character

import "fmt"

// Ability 인터페이스 정의
type Ability interface {
	Effect() string
	Use(c *Character, target *Character)
}

// SecondWind 구현체
type SecondWind struct{}

func (s SecondWind) Effect() string {
	return "Heals 10 HP."
}

func (s SecondWind) Use(c *Character, target *Character) {
	c.HP += 10
	if c.HP > 100 {
		c.HP = 100
	}
	fmt.Printf("%s used Second Wind and healed 10 HP!\n", c.Name)
}

// Fireball 구현체
type Fireball struct{}

func (f Fireball) Effect() string {
	return "Deals 15 damage to the target."
}

func (f Fireball) Use(c *Character, target *Character) {
	damage := 15 - target.Defense
	if damage < 0 {
		damage = 0
	}
	target.HP -= damage
	fmt.Printf("%s cast Fireball and dealt %d damage to %s!\n", c.Name, damage, target.Name)
}
