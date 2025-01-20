package character

type Character struct {
	Name       string
	Class      string
	Level      int
	HP         int
	Attack     int
	Defense    int
	Feature    Feature
	Ability    Ability
	IsDefeated bool
}
