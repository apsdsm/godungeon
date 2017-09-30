package game

// Navigatable returns positional information about an object that can be navigated by a Navigator.
type Navigatable interface {
	GetRelativeTile(tile *Tile, direction Direction) (*Tile, error)
}
