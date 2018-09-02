package controllers

import (
	"math"
	"math/rand"

	"github.com/apsdsm/godungeon/game"
)

// An ActorController contains methods for interacting with an actor object.
type ActorController struct {
	seed   int64
	random *rand.Rand
}

// ActorControllerConfig stores config data for controller
type ActorControllerConfig struct {
	Seed int64
}

// NewActorController return a new initialized actor controller
func NewActorController(config ActorControllerConfig) ActorController {
	a := ActorController{}

	a.seed = config.Seed
	a.random = rand.New(rand.NewSource(a.seed))

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

	// check if hits
	isHit := int(math.Floor(100*c.random.Float64())) < actor.Attack.ChanceToHit

	if !isHit {
		return nil
	}

	// pre casting to float makes the code below easier to read
	hitMin := float64(actor.Attack.MinDamage)
	hitMax := float64(actor.Attack.MaxDamage)
	cutRate := float64(target.Defence.DamageCutRate)
	cutCeil := float64(target.Defence.DamageCutCeil)

	// total raw damage
	hit := hitMin + math.Ceil((hitMax-hitMin)*c.random.Float64())

	// amount of damage deflected
	cut := math.Ceil(hit * 0.01 * cutRate)

	// make sure only a max amount can be deflected
	if cut > cutCeil {
		cut = cutCeil
	}

	// actual damage inflicted
	damage := int(hit - cut)

	// inflict damage
	target.Hp -= damage

	// check if target died this hit
	if target.Hp <= 0 {
		target.Hp = 0
		target.IsDead = true
	}

	return nil
}
