package controllers

import (
	"fmt"

	"github.com/apsdsm/godungeon/game"
)

// IllegalMove is an error that is returned when trying to performed a move that is logically impossible.
type IllegalMove struct {
	ToTile *game.Tile
}

// Error returns a string representation of the IllegalMove error.
func (e IllegalMove) Error() string {
	return fmt.Sprintf("cannot move to specified tile (%d, %d)", e.ToTile.Position.X, e.ToTile.Position.Y)
}
