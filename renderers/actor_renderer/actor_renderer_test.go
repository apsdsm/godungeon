package actor_renderer_test

import (
	"github.com/apsdsm/canvas"

	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/renderers/actor_renderer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActorRenderer", func() {

	Context("with nothing in entity array", func() {
		It("renders nothing", func() {
			entities := []game.Actor{}
			layer := canvas.NewLayer(4, 4, 0, 0)

			renderer := actor_renderer.New(&entities, &layer)
			renderer.Render()

			Expect(layer.Grid[1][1].Rune).To(Equal(int32(0)))
		})
	})
	Context("with actor in entity array", func() {
		It("renders actor to the layer", func() {

			tile := game.Tile{
				Position: game.Position{1, 1},
			}

			actor := game.Actor{
				Tile:       &tile,
				Appearance: 'x',
			}

			entities := []game.Actor{
				actor,
			}

			layer := canvas.NewLayer(4, 4, 0, 0)

			renderer := actor_renderer.New(&entities, &layer)
			renderer.Render()

			Expect(layer.Grid[1][1].Rune).To(Equal('x'))
		})
	})
})
