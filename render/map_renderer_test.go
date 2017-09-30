package render_test

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/io"
	"github.com/apsdsm/godungeon/render"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapRenderer", func() {
	It("passes map data to a layer", func() {

		layer := canvas.NewLayer(4, 4, 0, 0)
		dungeon := io.LoadMap("../fixtures/maps/small.json")
		renderer := render.NewMapRenderer()

		renderer.DrawMap(&dungeon, layer)

		Expect(layer.Grid[0][0].Rune).To(Equal('#'))
	})
})
