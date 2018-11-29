package updaters

import (
	"github.com/apsdsm/godungeon/controllers"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
)

// A Player is used to convert user input actions to inputHandler that happen on the player's actor object
type Player struct {
	actor           *game.Actor
	actorController *controllers.ActorController
	inputHandler    *input.TcellInputHandler
	boundMovements  map[input.Key]game.Direction
}

// A PlayerConfig transports setup information to new Player instances
type PlayerConfig struct {
	Actor           *game.Actor
	ActorController *controllers.ActorController
	Input           *input.TcellInputHandler
}

// NewPlayer creates and returns a new player updater
func NewPlayer(config PlayerConfig) Player {
	p := Player{}
	p.actor = config.Actor
	p.actorController = config.ActorController
	p.inputHandler = config.Input
	p.boundMovements = make(map[input.Key]game.Direction)
	return p
}

// BindMovement binds a key to a movement direction.
func (p *Player) BindMovement(key input.Key, direction game.Direction) {
	p.boundMovements[key] = direction
}

// Update will check for input events, and carry out the player's intended action appropriately.
func (p *Player) Update() {
	events := p.inputHandler.Events()

	for _, key := range events.Keys {
		if boundMovement, ok := p.boundMovements[key]; ok {

			// @fixme - this could crash if neighbor is nil
			tile := p.actor.Tile.Neighbor(boundMovement)

			if tile.Occupant != nil {
				p.actorController.Attack(p.actor, tile.Occupant)
				continue
			}

			if tile.Walkable {
				p.actorController.Move(p.actor, tile)
				continue
			}
		}
	}
}
