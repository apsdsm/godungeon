package controllers

import (
	"fmt"

	"github.com/apsdsm/godungeon/game"
	"github.com/go-errors/errors"
)

// An ActorController contains methods for interacting with an actor object.
type ActorController struct{}

// IllegalMove is an error that is returned when trying to move an actor somewhere
// they can't go. It contains a pointer to the tile.
type IllegalMove struct {
	ToTile *game.Tile
}

// Error returns a string representation of the IllegalMove error.
func (e IllegalMove) Error() string {
	return fmt.Sprintf("cannot move to specified tile (%d, %d)", e.ToTile.Position.X, e.ToTile.Position.Y)
}

// Move will change the tile position for an actor. If the tile cannot be moved to
// because it is occupied or not walkable, the function returns an IllegalMove error.
func (c *ActorController) Move(actor *game.Actor, tile *game.Tile) error {
	if !tile.Walkable || tile.Occupant != nil {
		return &IllegalMove{tile}
	}

	actor.Tile.Occupant = nil
	tile.Occupant = actor
	actor.Tile = tile
	return nil
}

// @todo implement
func (c *ActorController) Attack(actor *game.Actor, target *game.Actor) error {
	return errors.New("not implemented")
}
