package game

// A NavigationHandler allows an entity to move to free tiles in a dungeon
type NavigationHandler struct {
	navigatable Navigatable
	actor       Actor
	currentTile *Tile
}

// NewNavigationHandler creates a new instance of NavigationHandler
func NewNavigationHandler(actor Actor, navigatable Navigatable) *NavigationHandler {
	d := NavigationHandler{}
	d.actor = actor
	d.navigatable = navigatable

	return &d
}

// Moves tries to move the navigator in the specified direction, occupying the tile if possible.
func (d *NavigationHandler) Move(direction Direction) {
	tile, _ := d.navigatable.GetRelativeTile(d.currentTile, direction)

	if tile == nil {
		return
	}

	if tile.Walkable && tile.Occupant == nil {
		tile.Occupant = d.actor
		d.currentTile.Occupant = nil
		d.currentTile = tile
	}
}

// Occupy directly moves the navigator to the specified tile if possible
func (d *NavigationHandler) OccupyTile(tile *Tile) {
	tile.Occupant = d.actor
	d.currentTile = tile
}
