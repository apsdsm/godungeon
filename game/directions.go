package game

// Direction is a generic type.
type Direction int

// Directions can be specified using their full names.
const (
	North = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

// Directions can be specified using their initials.
const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)
