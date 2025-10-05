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

var _ = Describe("drawing lines", func() {

	var (
		style tcell.Style
		layer canvas.Layer
	)

	BeforeEach(func() {
		style = tcell.StyleDefault
		layer = canvas.NewLayer(10, 10, 0, 0)
	})

	// draws -> ────── (x10)
	It("draws a horizonal line", func() {
		painter.DrawHLine(&layer, 0, 0, 10, style)

		for i := 0; i < 10; i++ {
			Expect(layer.Grid[i][0].Rune).To(Equal('─'))
		}
	})

	//          │
	// draws -> │
	//          │ (x10)
	It("draws a vertical line", func() {
		painter.DrawVLine(&layer, 0, 0, 10, style)

		for i := 0; i < 10; i++ {
			Expect(layer.Grid[0][i].Rune).To(Equal('│'))
		}
	})
})
