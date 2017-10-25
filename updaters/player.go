package updaters

import (
	"github.com/apsdsm/godungeon/controllers"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
)

// A Player is used to convert user input actions to inputHandler that happen on the player's actor object
type Player struct {
	actor           *game.Actor
	inputHandler    InputHandlerForPlayer
	boundMovements  map[input.Key]game.Direction
	ActorController ActorControllerForPlayer
}

// @todo should this be defined closer to the actor controller? or just make a player package?
// ActorControllerForPlayer defines the commands the player will give to actors
type ActorControllerForPlayer interface {
	Move(actor *game.Actor, tile *game.Tile) error
	Attack(actor *game.Actor, target *game.Actor) error
}

// @todo should this be defined closer to the input handler? or just make a player package?
// InputHandlerForPlayer defines the iterface for an input handler
type InputHandlerForPlayer interface {
	Events() input.Events
}

// NewController creates and returns a new controller
func NewPlayer(actor *game.Actor, inputHandler InputHandlerForPlayer) Player {
	c := Player{}
	c.actor = actor
	c.inputHandler = inputHandler
	c.boundMovements = make(map[input.Key]game.Direction)
	c.ActorController = &controllers.ActorController{}

	return c
}

// BindMovement binds a key to a movement direction. When the key is pressed, the actor receives a Move request in the
// bound direction.
func (c *Player) BindMovement(key input.Key, direction game.Direction) {
	c.boundMovements[key] = direction
}

// Handle will handle input in the given context, and return a new context containing any changes made by the input
func (c *Player) Update() {
	events := c.inputHandler.Events()

	for _, key := range events.Keys {
		if boundMovement, ok := c.boundMovements[key]; ok {
			tile := c.actor.Tile.Neighbor(boundMovement)

			if tile.Occupant != nil {
				c.ActorController.Attack(c.actor, tile.Occupant)
				continue
			}

			if tile.Walkable {
				c.ActorController.Move(c.actor, tile)
				continue
			}
		}
	}
}
