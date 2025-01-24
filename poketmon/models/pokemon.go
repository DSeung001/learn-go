package models

// Pokemon represents the attributes of a Pokemon.
type Pokemon struct {
	Level        int
	Name         string
	Type         string
	Evolution    string
	PreEvolution string
	Stats        map[string]int
	Moves        []Move
}
