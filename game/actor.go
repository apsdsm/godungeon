package game

// An Actor accepts commands from a controller (or various controllers) and passes that on to the appropriate handler
// object.
type Actor interface {
	Move(direction Direction)
	//	MoveToTile(tile *Tile) error
	//	GetRelativeTile(direction Direction) (*Tile, error)
}
