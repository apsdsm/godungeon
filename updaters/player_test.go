package updaters

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apsdsm/godungeon/fakes"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
)

var _ = Describe("Player", func() {

	var (
		player          Player
		actorController *fakes.FakeActorController
		inputHandler    *fakes.FakeInputHandler
		fromTile        *game.Tile
		toTile          *game.Tile
		actor           *game.Actor
		occupant        *game.Actor
	)

	BeforeEach(func() {
		inputHandler = &fakes.FakeInputHandler{}
		actorController = &fakes.FakeActorController{}

		fromTile = &game.Tile{}
		toTile = &game.Tile{}
		actor = &game.Actor{}
		occupant = &game.Actor{}
		actorController = &fakes.FakeActorController{}

		player = NewPlayer(actor, inputHandler)
		player.ActorController = actorController
		player.BindMovement(input.NewKey(input.KeyRune, 'j'), game.West)
	})

	Context("updating", func() {
		It("get input from input handler each update", func() {
			player.Update()
			Expect(inputHandler.Received("Events"))
		})
	})

	Context("navigating", func() {
		Context("when trying to move to empty tile", func() {
			It("moves actor to tile", func() {
				fromTile.Occupant = actor
				fromTile.Walkable = true
				fromTile.Neighbors[game.W] = toTile

				toTile.Occupant = nil
				toTile.Walkable = true

				actor.Tile = fromTile

				inputHandler.GetEventsRet = makeKeyEvent(input.KeyRune, 'j')

				player.Update()

				Expect(actorController.Received("Move", fakes.MoveSig{actor, toTile}))
			})
		})

		Context("when trying to move to an impassible tile", func() {
			It("does not move", func() {
				fromTile.Occupant = actor
				fromTile.Walkable = true
				fromTile.Neighbors[game.W] = toTile

				toTile.Occupant = occupant
				toTile.Walkable = false

				actor.Tile = fromTile

				inputHandler.GetEventsRet = makeKeyEvent(input.KeyRune, 'j')

				player.Update()

				Expect(actorController.DidNotReceive("Move"))
			})
		})

		Context("when trying to move to an inhabited tile", func() {
			It("attacks the inhabitant, without moving to tile", func() {
				fromTile.Occupant = actor
				fromTile.Walkable = true
				fromTile.Neighbors[game.W] = toTile

				toTile.Occupant = occupant
				toTile.Walkable = true

				actor.Tile = fromTile
				occupant.Tile = toTile

				inputHandler.GetEventsRet = makeKeyEvent(input.KeyRune, 'j')

				player.Update()

				Expect(actorController.Received("Attack", fakes.AttackSig{actor, occupant}))
			})
		})
	})
})

func makeKeyEvent(key input.KeyCode, ch rune) input.Events {
	e := input.Events{}
	e.Keys = make([]input.Key, 1, 1)
	e.Keys[0] = input.NewKey(key, ch)

	return e
}
