package game

type Actor struct {
	Name     string
	Link     string
	Rune     rune
	Tile     *Tile
	Attack   Attack
	IsPlayer bool
	HP       int
	MP       int
}
