package character

// Feature 인터페이스 정의
type Feature interface {
	Description() string
	ApplyEffect(c *Character)
}

// HeavyArmorProficiency 구현체
type HeavyArmorProficiency struct{}

func (h HeavyArmorProficiency) Description() string {
	return "Reduces incoming damage by 2."
}

func (h HeavyArmorProficiency) ApplyEffect(c *Character) {
	c.Defense += 2
}

// Spellcasting 구현체
type Spellcasting struct{}

func (s Spellcasting) Description() string {
	return "Increases attack power by 3."
}

func (s Spellcasting) ApplyEffect(c *Character) {
	c.Attack += 3
}
