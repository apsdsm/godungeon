package painter_test

import (
	"github.com/apsdsm/canvas/painter"
	"github.com/gdamore/tcell"

	"github.com/apsdsm/canvas"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fill", func() {

	var (
		style tcell.Style
		layer canvas.Layer
	)

	BeforeEach(func() {
		style = tcell.StyleDefault
		layer = canvas.NewLayer(10, 10, 0, 0)
	})

	It("fills the layer with a single rune", func() {
		painter.Fill(&layer, '!', style)

		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				Expect(layer.At(x, y).Rune).To(Equal('!'))
			}
		}
	})

	It("does not paint a double width rune if there isn't enough space", func() {
		painter.Fill(&layer, 'の', style)

		Expect(layer.Grid[9][9].Rune).To(Equal(' '))
	})
})
