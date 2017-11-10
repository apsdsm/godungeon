package controllers

import (
	"errors"

	"github.com/apsdsm/godungeon/game"
)

// An ActorController contains methods for interacting with an actor object.
type ActorController struct {
	damageCalculator game.DamageCalculator
}

// NewActorController return a new initialized actor controller
func NewActorController() ActorController {
	a := ActorController{}
	a.damageCalculator = game.NewDamageCalculator()
	return a
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

	// it should get the attack from one actor, and calc damage, then apply damage to target
	// damage =
	// 1. check if hit or miss
	// 2. calc attack max damage (random between min and max damage)
	// 3. minus target defense
	// if damage < 0, damage = 0
}
