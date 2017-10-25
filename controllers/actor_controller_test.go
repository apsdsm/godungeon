package controllers_test

import (
	. "github.com/apsdsm/godungeon/controllers"

	"github.com/apsdsm/godungeon/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActorController", func() {
	Describe("Move", func() {
		Context("provided with unoccupied tile", func() {
			var (
				actor    game.Actor
				fromTile game.Tile
				toTile   game.Tile
				err      error
			)

			BeforeEach(func() {
				actor = game.Actor{
					Tile: &fromTile,
				}
				fromTile = game.Tile{
					Occupant: &actor,
					Walkable: true,
				}
				toTile = game.Tile{
					Walkable: true,
				}

				controller := ActorController{}
				err = controller.Move(&actor, &toTile)
			})

			It("moves actor to tile", func() {
				Expect(toTile.Occupant).To(Equal(&actor))
				Expect(fromTile.Occupant).To(BeNil())
			})

			It("provides a nil error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("provided with occupied tile", func() {
			var (
				actor    game.Actor
				occupant game.Actor
				fromTile game.Tile
				toTile   game.Tile
				err      error
			)

			BeforeEach(func() {
				actor = game.Actor{
					Tile: &fromTile,
				}
				occupant = game.Actor{
					Tile: &toTile,
				}
				fromTile = game.Tile{
					Occupant: &actor,
					Walkable: true,
				}
				toTile = game.Tile{
					Walkable: true,
					Occupant: &occupant,
				}

				controller := ActorController{}
				err = controller.Move(&actor, &toTile)
			})

			It("does not move actor to tile", func() {
				Expect(toTile.Occupant).To(Equal(&occupant))
				Expect(fromTile.Occupant).To(Equal(&actor))
			})

			It("returns an IllegalMove error", func() {
				Expect(err).To(BeAssignableToTypeOf(&IllegalMove{}))
			})
		})

		Context("provided with unwalkable tile", func() {
			var (
				actor    game.Actor
				fromTile game.Tile
				toTile   game.Tile
				err      error
			)

			BeforeEach(func() {
				actor = game.Actor{
					Tile: &fromTile,
				}
				fromTile = game.Tile{
					Occupant: &actor,
					Walkable: true,
				}
				toTile = game.Tile{
					Walkable: false,
				}

				controller := ActorController{}
				err = controller.Move(&actor, &toTile)
			})

			It("does not move actor to tile", func() {
				Expect(toTile.Occupant).To(BeNil())
				Expect(fromTile.Occupant).To(Equal(&actor))
			})

			It("returns an IllegalMove error", func() {
				Expect(err).To(BeAssignableToTypeOf(&IllegalMove{}))
			})

		})
	})
})
