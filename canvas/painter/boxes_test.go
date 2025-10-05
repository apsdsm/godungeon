package painter_test

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/canvas/painter"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boxes", func() {
	var (
		style tcell.Style
		layer canvas.Layer
	)

	BeforeEach(func() {
		style = tcell.StyleDefault
		layer = canvas.NewLayer(10, 10, 0, 0)
	})

	It("draws a box", func() {
		painter.DrawBox(&layer, 0, 0, 5, 5, style)

		// expect corners to be drawn correctly
		Expect(layer.Grid[0][0].Rune).To(Equal('┌'))
		Expect(layer.Grid[5][0].Rune).To(Equal('┐'))
		Expect(layer.Grid[5][5].Rune).To(Equal('┘'))
		Expect(layer.Grid[0][5].Rune).To(Equal('└'))

		// expect top and bottom bar to be drawn correctly
		for i := 1; i < 4; i++ {
			Expect(layer.Grid[i][0].Rune).To(Equal('─'))
			Expect(layer.Grid[i][5].Rune).To(Equal('─'))
		}

		// expect left and right bar to be drawn correctly
		for i := 1; i < 4; i++ {
			Expect(layer.Grid[0][i].Rune).To(Equal('│'))
			Expect(layer.Grid[5][i].Rune).To(Equal('│'))
		}
	})

	It("does not draw a box out of bounds", func() {
		Expect(func() {
			painter.DrawBox(&layer, 0, 0, 20, 20, style)
		}).ShouldNot(Panic())
	})
})
