package combat

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func NewMonster(name string, hp int, attack int) *Monster {
	return &Monster{
		Name:      name,
		MaxHP:     hp,
		CurrentHP: hp,
		Attack:    attack,
	}
}