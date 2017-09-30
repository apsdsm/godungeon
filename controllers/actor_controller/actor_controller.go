package actor_controller

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
)

// A ActorController is used to convert user input actions to events that happen on the actor_controller object
type ActorController struct {
	actor game.Actor
	state State

	boundMovements map[input.Key]game.Direction
}

type State int

const (
	Navigating = iota
)

// NewController creates and returns a new controller
func New(player game.Actor) *ActorController {
	c := ActorController{}
	c.actor = player
	c.boundMovements = make(map[input.Key]game.Direction)

	return &c
}

// SetState changes the current controller state
func (c *ActorController) SetState(state State) {
	c.state = state
}

// BindMovement binds a key to a movement direction. When the key is pressed, the actor receives a Move request in the
// bound direction.
func (c *ActorController) BindMovement(key input.Key, direction game.Direction) {
	c.boundMovements[key] = direction
}

// Handle will handle input in the given context, and return a new context containing any changes made by the input
func (c *ActorController) Update(events input.Events) {
	switch c.state {
	case Navigating:
		c.updateNavigating(&events)
	}
}

// updateNavigating handles input while the controller is navigating
func (c *ActorController) updateNavigating(events *input.Events) {
	for _, key := range events.Keys {
		if boundMovement, ok := c.boundMovements[key]; ok {
			c.actor.Move(boundMovement)
		}
	}
}
