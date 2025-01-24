package models

type Pokemon struct {
	Level        int
	Name         string
	Type         string
	Evolution    string
	PreEvolution string
	Stats        Stats
	Moves        []Move
	Live         bool
	Trainer      string
}

type Stats struct {
	CurrentHp int
	Hp        int
	Attack    int
	Defense   int
	SpAtk     int
	SpDef     int
	Speed     int
}
