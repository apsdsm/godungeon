package dungeon_renderer_test

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/file"
	"github.com/apsdsm/godungeon/renderers/dungeon_renderer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DungeonRenderer", func() {
	It("renders a map to a layer", func() {

		layer := canvas.NewLayer(3, 4, 0, 0)
		dungeon := file.LoadMap("../../fixtures/maps/small.json")

		renderer := dungeon_renderer.New(dungeon, &layer)

		renderer.Render()

		Expect(string(layer.Grid[0][0].Rune)).To(Equal(string('#')))
	})
})
