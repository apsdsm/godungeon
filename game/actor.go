package game

type Actor struct {
	Name     string
	Link     string
	Rune     rune
	Tile     *Tile
	Attack   Attack
	Defence  Defence
	IsPlayer bool
	IsDead   bool
	Hp       int
	MaxHp    int
	Mp       int
	MaxMp    int
}

// HpPercentRemaining returns the remaining HP for this actor as a decimal percent
func (a *Actor) HpPercentRemaining() float64 {
	return float64(a.Hp) / float64(a.MaxHp)
}
