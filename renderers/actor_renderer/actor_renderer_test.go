package actor_renderer_test

import (
	"github.com/apsdsm/canvas"

	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/renderers/actor_renderer"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActorRenderer", func() {

	var (
		renderer *actor_renderer.ActorRenderer
	)

	Context("with nothing in entity array", func() {
		var (
			entities []game.Actor
			layer    canvas.Layer
		)

		BeforeEach(func() {
			entities = []game.Actor{}
			layer = canvas.NewLayer(1, 1, 0, 0)
			renderer = actor_renderer.New(&entities, &layer)
		})

		It("renders nothing", func() {
			renderer.Render()
			Expect(layer.Grid[0][0].Rune).To(Equal(int32(0)))
		})
	})

	Context("with actor in entity array", func() {

		var (
			tiles     []game.Tile
			actor     game.Actor
			other     game.Actor
			lowHealth game.Actor
			entities  []game.Actor
			layer     canvas.Layer
		)

		BeforeEach(func() {
			tiles = []game.Tile{
				{
					Position: game.NewPosition(0, 0),
				},
				{
					Position: game.NewPosition(1, 0),
				},
				{
					Position: game.NewPosition(2, 0),
				},
			}

			actor = game.Actor{
				Tile:     &tiles[0],
				Rune:     'x',
				IsPlayer: true,
			}

			other = game.Actor{
				Tile:     &tiles[1],
				Rune:     'y',
				Hp:       10,
				MaxHp:    10,
				IsPlayer: false,
			}

			lowHealth = game.Actor{
				Tile:     &tiles[2],
				Rune:     'z',
				Hp:       1,
				MaxHp:    10,
				IsPlayer: false,
			}

			entities = []game.Actor{
				actor,
				other,
				lowHealth,
			}

			layer = canvas.NewLayer(3, 1, 0, 0)

			renderer = actor_renderer.New(&entities, &layer)
		})

		It("renders player to the layer", func() {
			renderer.Render()
			Expect(layer.Grid[0][0].Rune).To(Equal('x'))
			Expect(layer.Grid[0][0].Style).To(Equal(tcell.StyleDefault.Foreground(game.White)))
		})

		It("renders mob to the layer", func() {
			renderer.Render()
			Expect(layer.Grid[1][0].Rune).To(Equal('y'))
			Expect(layer.Grid[1][0].Style).To(Equal(tcell.StyleDefault.Foreground(game.Green)))
		})

		It("renders low health mob to laeyr", func() {
			renderer.Render()
			Expect(layer.Grid[2][0].Rune).To(Equal('z'))
			Expect(layer.Grid[2][0].Style).To(Equal(tcell.StyleDefault.Foreground(game.Red)))
		})
	})
})
