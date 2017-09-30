package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apsdsm/godungeon/fakes"
	"github.com/apsdsm/godungeon/game"
)

var _ = Describe("DungeonController", func() {

	It("moves to occupy a free tile", func() {
		actor := fakes.NewActor()

		fromTile := &game.Tile{
			Walkable: true,
		}

		toTile := &game.Tile{
			Walkable: true,
		}

		navigatable := fakes.NewNavigatable()
		navigatable.GetRelativeTileReturns = fakes.GetRelativeTileRet{toTile, nil}

		navigator := game.NewNavigationHandler(actor, navigatable)
		navigator.OccupyTile(fromTile)

		// move to new tile
		navigator.Move(game.North)

		// checked to see if correct tile was available
		Expect(navigatable.Received("GetRelativeTile", fakes.GetRelativeTileSig{fromTile, game.North})).To(BeTrue())

		// moved to new tile
		Expect(toTile.Occupant).To(Equal(actor))

		// abandoned prev tile
		Expect(fromTile.Occupant).To(BeNil())

	})

	It("does not move to impassible tile", func() {
		actor := fakes.NewActor()

		fromTile := &game.Tile{
			Walkable: true,
		}

		toTile := &game.Tile{
			Walkable: false,
		}

		navigatable := fakes.NewNavigatable()
		navigatable.GetRelativeTileReturns = fakes.GetRelativeTileRet{toTile, nil}

		navigator := game.NewNavigationHandler(actor, navigatable)
		navigator.OccupyTile(fromTile)

		// move to new tile
		navigator.Move(game.North)

		// checked to see if correct tile was available
		Expect(navigatable.Received("GetRelativeTile", fakes.GetRelativeTileSig{fromTile, game.North})).To(BeTrue())

		// moved to new tile
		Expect(toTile.Occupant).To(BeNil())

		// abandoned prev tile
		Expect(fromTile.Occupant).To(Equal(actor))
	})

	It("does not move to occupied tile", func() {
		actor := fakes.NewActor()
		occupant := fakes.NewActor()

		fromTile := &game.Tile{
			Walkable: true,
		}

		toTile := &game.Tile{
			Walkable: true,
			Occupant: occupant,
		}

		navigatable := fakes.NewNavigatable()
		navigatable.GetRelativeTileReturns = fakes.GetRelativeTileRet{toTile, nil}

		navigator := game.NewNavigationHandler(actor, navigatable)
		navigator.OccupyTile(fromTile)

		// move to new tile
		navigator.Move(game.North)

		// checked to see if correct tile was available
		Expect(navigatable.Received("GetRelativeTile", fakes.GetRelativeTileSig{fromTile, game.North})).To(BeTrue())

		// moved to new tile
		Expect(toTile.Occupant).To(Equal(occupant))

		// abandoned prev tile
		Expect(fromTile.Occupant).To(Equal(actor))
	})

	It("does not move if no tile present", func() {
		actor := fakes.NewActor()

		fromTile := &game.Tile{
			Walkable: true,
			Occupant: actor,
		}

		navigatable := fakes.NewNavigatable()
		navigatable.GetRelativeTileReturns = fakes.GetRelativeTileRet{nil, nil}

		navigator := game.NewNavigationHandler(actor, navigatable)
		navigator.OccupyTile(fromTile)

		// move to new tile
		navigator.Move(game.North)

		// checked to see if correct tile was available
		Expect(navigatable.Received("GetRelativeTile", fakes.GetRelativeTileSig{fromTile, game.North})).To(BeTrue())

		// abandoned prev tile
		Expect(fromTile.Occupant).To(Equal(actor))
	})

})
