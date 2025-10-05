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

package fakes

import (
	"bytes"

	"github.com/gdamore/tcell"
)

type cell struct {
	codePoint rune
	fgColor   tcell.Color
	bgColor   tcell.Color
}

// ScreenBridge is a low level fake screen bridge for testing. It does not draw to the screen,
// but keeps track of everything that is drawn inside of it. It can be queried to return
// values for previous draw methods.
type ScreenBridge struct {
	width  int
	height int

	cells [][]cell

	CalledSetContent int
	CalledSize       int
	CalledShow       int
}

// NewScreenBridge returns a new ScreenBridge object. The object is initialized with the
// specified width and height.
func NewScreenBridge(width, height int) *ScreenBridge {
	s := ScreenBridge{
		width:  width,
		height: height,
	}

	s.cells = make([][]cell, width)

	for i := range s.cells {
		s.cells[i] = make([]cell, height)
	}

	return &s
}

// SetContent is a dummy method for tcell's SetContent (https://godoc.org/github.com/gdamore/tcell#CellBuffer.SetContent)
func (f *ScreenBridge) SetContent(x int, y int, mainc rune, combc []rune, style tcell.Style) {
	f.CalledSetContent++
	f.cells[x][y].codePoint = mainc
	f.cells[x][y].fgColor, f.cells[x][y].bgColor, _ = style.Decompose()
}

// Size is a dummy method for tcel's Size (https://godoc.org/github.com/gdamore/tcell#CellBuffer.Size)
func (f *ScreenBridge) Size() (int, int) {
	f.CalledSize++
	return f.width, f.height
}

// Show is a dummy method for tcel's Show (TODO: link here!!)
func (f *ScreenBridge) Show() {
	f.CalledShow++
}

// GetLine returns the input that was generated for line at y
func (f *ScreenBridge) GetLine(y, x, len int) string {
	var buffer bytes.Buffer

	for i := 0; i < len; i++ {

		if x >= f.width {
			break
		}

		// get rune
		current := f.cells[x][y].codePoint

		// write rune to buffer
		_, _ = buffer.WriteRune(current)

		// advance x
		x++
	}

	return buffer.String()
}

// GetStyleAt returns the tcell style used in a cell
func (f *ScreenBridge) GetStyleAt(x, y int) (fg, bg tcell.Color) {
	return f.cells[x][y].fgColor, f.cells[x][y].bgColor
}

// GetRuneAt returns the rune used in the cell
func (f *ScreenBridge) GetRuneAt(x, y int) (r rune) {
	return f.cells[x][y].codePoint
}
