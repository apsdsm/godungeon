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

package canvas

import "github.com/gdamore/tcell"

// A Layer is a drawable area in a canvas that carries its own origin coords
// and max boundaries.
type Layer struct {
	Grid                            [][]Cell
	Width, Height, X, Y, MaxX, MaxY int
}

// NewLayer creates a new layer with the specified size at the specified offset
func NewLayer(width, height, x, y int) Layer {
	l := Layer{}

	l.Grid = make([][]Cell, width)

	for i := range l.Grid {
		l.Grid[i] = make([]Cell, height)
		for y := range l.Grid[i] {
			l.Grid[i][y].Rune = 0
			l.Grid[i][y].Style = tcell.StyleDefault
		}
	}

	l.Width = width
	l.Height = height
	l.X = x
	l.Y = y
	l.MaxX = width - 1
	l.MaxY = height - 1

	return l
}

// SetRune will set the Cell at the given coordinates with the provided rune.
func (l *Layer) SetRune(x, y int, r rune) {
	if l.InBounds(x, y) {
		l.Grid[x][y].Rune = r
	}
}

func (l *Layer) SetStyle(x, y int, style tcell.Style) {
	if l.InBounds(x, y) {
		l.Grid[x][y].Style = style
	}
}

// Normalize will take a set of rect coords and normalize them to be legal layer coordinates.
func (l *Layer) Normalize(minX, minY, maxX, maxY int) (nMinX, nMinY, nMaxX, nMaxY int) {
	nMinX = normalize(minX, 0, l.MaxX)
	nMaxX = normalize(maxX, 0, l.MaxX)
	nMinY = normalize(minY, 0, l.MaxY)
	nMaxY = normalize(maxY, 0, l.MaxY)

	return
}

// normalize a number so that it is inside a range of a min and max.
func normalize(target, min, max int) int {
	if target < min {
		return min
	} else if target > max {
		return max
	} else {
		return target
	}
}

// InBounds returns true if cell is inside layer bounds
func (l *Layer) InBounds(x, y int) bool {
	return x >= 0 && x < l.Width && y >= 0 && y < l.Height
}

// At returns the cell at the given coords, or nil if no such position exists
func (l *Layer) At(x, y int) *Cell {
	if l.InBounds(x, y) {
		return &l.Grid[x][y]
	}
	return nil
}

// Each will pass each cell in the grid to the provided closure
func (l *Layer) Each(f func(*Cell)) {
	for x := range l.Grid {
		for y := range l.Grid[x] {
			f(&l.Grid[x][y])
		}
	}
}

// @todo write test
// Clear will set each of the runes in the layer to zero and the style to the default
func (l *Layer) Clear() {
	for x := range l.Grid {
		for y := range l.Grid[x] {
			l.Grid[x][y].Rune = 0
			l.Grid[x][y].Style = tcell.StyleDefault
		}
	}
}
