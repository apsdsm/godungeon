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

package painter_test

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/canvas/painter"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("drawing wrapped text", func() {

	var (
		style tcell.Style
		layer canvas.Layer
	)

	BeforeEach(func() {
		style = tcell.StyleDefault
		layer = canvas.NewLayer(10, 10, 0, 0)
	})

	It("paints over a section of the screen with a single rune", func() {
		painter.Paint(&layer, 0, 0, 2, 2, '!', style)

		for x := 0; x <= 2; x++ {
			for y := 0; y <= 2; y++ {
				Expect(layer.Grid[x][y].Rune).To(Equal('!'))
			}
		}
	})

	It("does not paint a double width rune if there isn't enough space", func() {
		painter.Paint(&layer, 9, 9, 9, 9, 'の', style)

		Expect(layer.Grid[9][9].Rune).To(Equal(' '))
	})

	It("does not paint out of bounds", func() {
		Expect(func() {
			painter.Paint(&layer, 5, 5, 10, 10, '!', style)
		}).ShouldNot(Panic())
	})
})
