package game

// A SharedPoint is a position on the map shared by more than one tile.
type SharedPoint struct {
	point Vec2
	tiles []Tile
}
