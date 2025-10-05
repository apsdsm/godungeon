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

	It("writes text with wrapping inside a designated area", func() {
		painter.DrawWrappedText(&layer, 0, 0, 10, 10, "this is some text", style)

		Expect(getLayerLine(&layer, 0, 0, 7)).To(Equal("this is"))
		Expect(getLayerLine(&layer, 0, 1, 9)).To(Equal("some text"))
	})

	It("wraps text when eol start of new word", func() {
		painter.DrawWrappedText(&layer, 0, 0, 8, 8, "abc def ghi jkl", style)

		Expect(getLayerLine(&layer, 0, 0, 7)).To(Equal("abc def"))
		Expect(getLayerLine(&layer, 0, 1, 7)).To(Equal("ghi jkl"))
	})

	It("wraps text when eol is space", func() {
		painter.DrawWrappedText(&layer, 0, 0, 7, 7, "abc def ghi jkl", style)

		Expect(getLayerLine(&layer, 0, 0, 7)).To(Equal("abc def"))
		Expect(getLayerLine(&layer, 0, 1, 7)).To(Equal("ghi jkl"))
	})

	It("wraps very long text", func() {
		painter.DrawWrappedText(&layer, 0, 0, 6, 6, "abcdefghijkl", style)

		Expect(getLayerLine(&layer, 0, 0, 6)).To(Equal("abcdef"))
		Expect(getLayerLine(&layer, 0, 1, 6)).To(Equal("ghijkl"))
	})

	It("does not write beyond specified area", func() {
		painter.DrawWrappedText(&layer, 0, 0, 6, 0, "abcdefghijkl", style)

		Expect(getLayerLine(&layer, 0, 0, 6)).To(Equal("abcdef"))
		Expect(getLayerLine(&layer, 0, 1, 6)).To(Equal(""))
	})

	It("wraps text with line breaks", func() {
		emptyLine := getLayerLine(&layer, 0, 0, 3)

		painter.DrawWrappedText(&layer, 0, 0, 6, 6, "abcd\n\nefgh\n\nijkl", style)

		Expect(getLayerLine(&layer, 0, 0, 4)).To(Equal("abcd"))
		Expect(getLayerLine(&layer, 0, 1, 4)).To(Equal(emptyLine))
		Expect(getLayerLine(&layer, 0, 2, 4)).To(Equal("efgh"))
		Expect(getLayerLine(&layer, 0, 3, 4)).To(Equal(emptyLine))
		Expect(getLayerLine(&layer, 0, 4, 4)).To(Equal("ijkl"))
	})
})
