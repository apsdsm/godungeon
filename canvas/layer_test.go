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

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Layer", func() {

	It("creates a new layer", func() {

		l := NewLayer(10, 5, 1, 2)

		Expect(len(l.Grid)).To(Equal(10))

		for i := range l.Grid {
			Expect(len(l.Grid[i])).To(Equal(5))
		}

		Expect(l.Width).To(Equal(10))
		Expect(l.Height).To(Equal(5))
		Expect(l.X).To(Equal(1))
		Expect(l.Y).To(Equal(2))
		Expect(l.MaxX).To(Equal(9))
		Expect(l.MaxY).To(Equal(4))
	})

	It("sets a rune on a grid coord", func() {
		l := NewLayer(10, 10, 0, 0)

		l.SetRune(0, 0, '!')

		Expect(l.Grid[0][0].Rune).To(Equal('!'))
	})

	It("does not set a rune outside of bounds", func() {
		Expect(func() {
			l := NewLayer(10, 10, 0, 0)
			l.SetRune(20, 20, '!')
		}).ShouldNot(Panic())
	})

	It("normalizes rect coords to layer min/max", func() {
		l := NewLayer(10, 10, 0, 0)

		minX, minY, maxX, maxY := l.Normalize(-5, -5, 20, 20)

		Expect(minX).To(Equal(0))
		Expect(minY).To(Equal(0))
		Expect(maxX).To(Equal(9))
		Expect(maxY).To(Equal(9))
	})
})
