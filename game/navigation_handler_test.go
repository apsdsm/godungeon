package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apsdsm/godungeon/fakes"
	"github.com/apsdsm/godungeon/game"
)

var _ = Describe("DungeonController", func() {

	It("moves to occupy a free tile", func() {

		fromTile := &game.Tile{
			Walkable: true,
		}

		toTile := &game.Tile{
			Walkable: true,
		}

		navigatable := fakes.NewNavigatable()
		navigatable.GetRelativeTileReturns = fakes.GetRelativeTileRet{toTile, nil}

		navigator := game.NewNavigationHandler(navigatable)
		navigator.OccupyTile(fromTile)

		// move to new tile
		navigator.Move(game.North)

		// checked to see if correct tile was available
		Expect(navigatable.Received("GetRelativeTile", fakes.GetRelativeTileSig{fromTile, game.North})).To(BeTrue())

		// moved to new tile
		Expect(toTile.Occupant).To(Equal(navigator))

		// abandoned prev tile
		Expect(fromTile.Occupant).To(BeNil())

	})

	It("does not move to impassible tile", func() {

	})

	It("does not move to occupied tile", func() {

	})

})

//func makeActorWithTile(tile *game.Tile) *fakes.FakeActor {
//	actor := fakes.NewActor()
//	actor.MoveToTileReturns = fakes.MoveToTileRet{nil}
//	actor.GetRelativeTileReturns = fakes.GetRelativeTileRet{tile, nil}
//
//	return actor
//}
//
//func makeKeyEvent(key input.KeyCode, ch rune) input.Events {
//	e := input.Events{}
//	e.Keys = make([]input.Key, 1, 1)
//	e.Keys[0] = input.NewKey(key, ch)
//
//	return e
//}
