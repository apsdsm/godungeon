package game

// A NavigationHandler allows an entity to move to free tiles in a dungeon
type NavigationHandler struct {
	navigatable Navigatable
}

// NewNavigationHandler creates a new instance of NavigationHandler
func NewNavigationHandler(navigatable Navigatable) *NavigationHandler {
	d := NavigationHandler{
		navigatable,
	}

	return &d
}

// Moves tries to move the navigator in the specified direction, occupying the tile if possible.
func (d *NavigationHandler) Move(direction Direction) {

}

// Occupy directly moves the navigator to the specified tile if possible
func (d *NavigationHandler) OccupyTile(tile *Tile) {

}
