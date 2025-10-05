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

	It("draws text", func() {

		var drawTextTests = []struct {
			x, y, stringSize int
			text             string
		}{
			{0, 0, 4, "ほげ"},
			{0, 1, 8, "alphabet"},
		}

		for _, test := range drawTextTests {
			painter.DrawText(&layer, test.x, test.y, test.text, style)

			line := getLayerLine(&layer, test.x, test.y, test.stringSize)

			Expect(line).To(Equal(test.text))
		}
	})
})
