package actor_controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apsdsm/godungeon/controllers/actor_controller"
	"github.com/apsdsm/godungeon/fakes"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
)

var _ = Describe("DungeonController", func() {

	var (
		tile       *game.Tile
		actor      *fakes.FakeActor
		controller *actor_controller.ActorController
	)

	BeforeEach(func() {
		tile = &game.Tile{}
		actor = makeActorWithTile(tile)
		controller = actor_controller.New(actor)
		controller.SetState(actor_controller.Navigating)
	})

	It("passes bound movement commands to actor", func() {
		key := input.NewKey(input.KeyRune, 'j')
		events := makeKeyEvent(input.KeyRune, 'j')
		controller.BindMovement(key, game.West)

		controller.Update(events)

		Expect(actor.Received("Move", fakes.MoveSig{game.West})).To(BeTrue())
	})
})

func makeActorWithTile(tile *game.Tile) *fakes.FakeActor {
	actor := fakes.NewActor()
	actor.MoveToTileReturns = fakes.MoveToTileRet{nil}
	actor.GetRelativeTileReturns = fakes.GetRelativeTileRet{tile, nil}

	return actor
}

func makeKeyEvent(key input.KeyCode, ch rune) input.Events {
	e := input.Events{}
	e.Keys = make([]input.Key, 1, 1)
	e.Keys[0] = input.NewKey(key, ch)

	return e
}
