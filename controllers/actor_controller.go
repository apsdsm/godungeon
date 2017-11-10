package controllers

import (
	"errors"

	"github.com/apsdsm/godungeon/game"
)

// An ActorController contains methods for interacting with an actor object.
type ActorController struct{}

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
