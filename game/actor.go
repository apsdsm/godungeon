package game

type Actor struct {
	Name       string
	Appearance rune
	Tile       *Tile
	Attack     Attack
}
