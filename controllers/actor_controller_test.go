package controllers_test

import (
	. "github.com/apsdsm/godungeon/controllers"

	"github.com/apsdsm/godungeon/fakes"
	"github.com/apsdsm/godungeon/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActorController", func() {

	var (
		damageCalculator fakes.FakeDamageCalculator
		controller       ActorController
	)

	BeforeEach(func() {
		damageCalculator = fakes.NewDamageCalculator()
		controller = NewActorController(&damageCalculator)
	})

	Describe("Move", func() {
		var (
			actor    game.Actor
			occupant game.Actor
			fromTile game.Tile
			toTile   game.Tile
		)

		Context("provided with unoccupied tile", func() {
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
			})

			It("moves actor to tile", func() {
				_ = controller.Move(&actor, &toTile)
				Expect(toTile.Occupant).To(Equal(&actor))
				Expect(fromTile.Occupant).To(BeNil())
			})

			It("provides a nil error", func() {
				err := controller.Move(&actor, &toTile)
				Expect(err).To(BeNil())
			})
		})

		Context("provided with occupied tile", func() {
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
			})

			It("does not move actor to tile", func() {
				_ = controller.Move(&actor, &toTile)
				Expect(toTile.Occupant).To(Equal(&occupant))
				Expect(fromTile.Occupant).To(Equal(&actor))
			})

			It("returns an IllegalMove error", func() {
				err := controller.Move(&actor, &toTile)
				Expect(err).To(BeAssignableToTypeOf(&IllegalMove{}))
			})
		})

		Context("provided with unwalkable tile", func() {
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
			})

			It("does not move actor to tile", func() {
				_ = controller.Move(&actor, &toTile)
				Expect(toTile.Occupant).To(BeNil())
				Expect(fromTile.Occupant).To(Equal(&actor))
			})

			It("returns an IllegalMove error", func() {
				err := controller.Move(&actor, &toTile)
				Expect(err).To(BeAssignableToTypeOf(&IllegalMove{}))
			})
		})
	})

	// @todo implement the attack
	Describe("Attack", func() {
		var (
			actor   game.Actor
			target  game.Actor
			attack  game.Attack
			defence game.Defence
		)

		Context("on hit", func() {
			BeforeEach(func() {
				attack.MaxDamage = 10
				actor.Attack = attack

				defence.DamageCutCeil = 10
				target.Defence = defence
				target.HP = 100

				damageCalculator.CalcDamageRet = fakes.CalcDamageRet{
					Damage: game.Damage{
						Dp:   20,
						Type: game.DamageHit,
					},
				}
			})

			It("applies damage to target", func() {
				_ = controller.Attack(&actor, &target)
				Expect(damageCalculator.Received("CalcDamage", fakes.CalcDamageSig{attack, defence})).To(BeTrue())
				Expect(target.HP).To(Equal(80))
			})

			It("return nil error", func() {
				err := controller.Attack(&actor, &target)
				Expect(err).To(BeNil())
			})
		})

		Context("on miss", func() {
			BeforeEach(func() {
				attack.MaxDamage = 10
				actor.Attack = attack

				defence.DamageCutCeil = 10
				target.Defence = defence
				target.HP = 100

				damageCalculator.CalcDamageRet = fakes.CalcDamageRet{
					Damage: game.Damage{
						Dp:   0,
						Type: game.DamageMiss,
					},
				}
			})

			It("applies no damage to target", func() {
				_ = controller.Attack(&actor, &target)
				Expect(damageCalculator.Received("CalcDamage", fakes.CalcDamageSig{attack, defence})).To(BeTrue())
				Expect(target.HP).To(Equal(100))
			})

			It("return nil error", func() {
				err := controller.Attack(&actor, &target)
				Expect(err).To(BeNil())
			})
		})
	})
})
