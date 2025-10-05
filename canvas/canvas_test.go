//    Copyright 2016 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package canvas_test

import (
	. "github.com/apsdsm/canvas"
	"github.com/apsdsm/canvas/fakes"
	"github.com/apsdsm/canvas/painter"
	"github.com/gdamore/tcell"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Canvas", func() {

	var (
		screen *fakes.ScreenBridge
		style  tcell.Style
	)

	BeforeEach(func() {
		screen = fakes.NewScreenBridge(10, 10)
		style = tcell.StyleDefault
	})

	Context("making layers", func() {

		It("adds default layer to the canvas that is the same size as screen", func() {
			c := NewCanvas(screen)

			c.AddLayer()

			Expect(len(c.Layers)).To(Equal(1))
			Expect(len(c.Layers[0].Grid)).To(Equal(10))

			for i := range c.Layers[0].Grid {
				Expect(len(c.Layers[0].Grid[i])).To(Equal(10))
			}

			Expect(c.Layers[0].X).To(Equal(0))
			Expect(c.Layers[0].Y).To(Equal(0))
		})

		It("adds a premade layer to the canvas", func() {
			c := NewCanvas(screen)
			l := NewLayer(10, 5, 1, 2)

			c.AddLayer(&l)

			Expect(len(c.Layers)).To(Equal(1))
			Expect(c.Layers[0]).To(BeEquivalentTo(&l))
		})

		It("adds new layers to top of stack", func() {
			c := NewCanvas(screen)
			l1 := NewLayer(10, 10, 1, 0)
			l2 := NewLayer(10, 10, 2, 0)

			c.AddLayer(&l1)

			Expect(len(c.Layers)).To(Equal(1))
			Expect(c.Layers[0]).To(BeEquivalentTo(&l1))

			c.AddLayer(&l2)

			Expect(len(c.Layers)).To(Equal(2))
			Expect(c.Layers[0]).To(BeEquivalentTo(&l2))

		})
	})

	Context("drawing to screen", func() {

		It("draws nothing to screen if there are no layers", func() {
			c := NewCanvas(screen)

			c.Draw()

			Expect(screen.CalledSetContent).To(Equal(0))
		})

		It("draws a single layer to screen", func() {
			c := NewCanvas(screen)
			l := NewLayer(10, 10, 0, 0)

			painter.DrawText(&l, 0, 0, "foobar", style)

			c.AddLayer(&l)
			c.Draw()

			Expect(screen.GetLine(0, 0, 6)).To(Equal("foobar"))
		})

		It("draws a single off-set layer to screen", func() {
			c := NewCanvas(screen)
			l := NewLayer(10, 10, 5, 5)

			painter.DrawText(&l, 0, 0, "foobar", style)

			c.AddLayer(&l)
			c.Draw()

			Expect(screen.GetLine(5, 5, 6)).To(Equal("fooba"))
		})

		It("draws one layer over another layer", func() {
			c := NewCanvas(screen)
			l1 := NewLayer(10, 10, 0, 0)
			l2 := NewLayer(10, 10, 0, 0)

			painter.DrawText(&l1, 0, 0, "foobar", style)
			painter.DrawText(&l2, 0, 0, "abcdef", style)

			c.AddLayer(&l1)
			c.AddLayer(&l2)
			c.Draw()

			Expect(screen.GetLine(0, 0, 6)).To(Equal("abcdef"))
		})

		It("updates the screen after drawing", func() {
			c := NewCanvas(screen)

			c.Draw()

			Expect(screen.CalledShow).To(Equal(1))
		})
	})

})
