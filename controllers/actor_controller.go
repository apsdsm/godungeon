package controllers

import (
	"fmt"

	"github.com/apsdsm/godungeon/debug"
	"github.com/apsdsm/godungeon/game"
)

// damageController is a local interface describing what the damage calculator should do
type damageCalculator interface {
	CalcDamage(attack game.Attack, defence game.Defence) game.Damage
}

// An ActorController contains methods for interacting with an actor object.
type ActorController struct {
	damageCalculator damageCalculator
}

// NewActorController return a new initialized actor controller
func NewActorController(damageCalculator damageCalculator) ActorController {
	a := ActorController{}
	a.damageCalculator = damageCalculator

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

// Attack will apply damage from one actor's attack to another actor.
func (c *ActorController) Attack(actor *game.Actor, target *game.Actor) error {

	debug.Log(fmt.Sprint(target))

	damage := c.damageCalculator.CalcDamage(actor.Attack, target.Defence)
	target.Hp -= damage.Dp

	// check if target died this hit
	if target.Hp <= 0 {
		target.Hp = 0
		target.IsDead = true
	}

	return nil
}
