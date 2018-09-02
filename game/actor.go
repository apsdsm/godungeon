package game

// Actor is a mob or the player
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
	Sight    int
}
