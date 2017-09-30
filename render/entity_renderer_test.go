package render_test

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/render"

	"github.com/apsdsm/godungeon/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EntityRenderer", func() {
	It("renders the actor to the screen", func() {

		player := game.Player{
			Rune: 'v',
			CurrentPosition: game.Position{
				X: 1,
				Y: 1,
			},
		}

		layer := canvas.NewLayer(4, 4, 0, 0)
		renderer := render.NewEntityRenderer()

		renderer.DrawPlayer(&player, layer)

		Expect(layer.Grid[1][1].Rune).To(Equal('v'))
	})
})
