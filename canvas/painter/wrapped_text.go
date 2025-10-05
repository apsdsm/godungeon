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

package painter

import (
	"bufio"
	"bytes"
	"unicode/utf8"

	"github.com/apsdsm/canvas"
	"github.com/gdamore/tcell"
)

// DrawWrappedText will print text to the screen, wrapping where possible to stay inside the bounds supplied
// in the paramters, and without overstepping the bounds of the screen.
func DrawWrappedText(layer *canvas.Layer, minX, minY, maxX, maxY int, text string, style tcell.Style) {

	if maxX > layer.MaxX {
		maxX = layer.MaxX
	}

	if maxY > layer.MaxY {
		maxY = layer.MaxY
	}

	if minX < 0 {
		minX = 0
	}

	if minY < 0 {
		minY = 0
	}

	x := minX
	y := minY
	textBuffer := bytes.NewBufferString(text)
	lineScanner := bufio.NewScanner(textBuffer)

	lineScanner.Split(bufio.ScanLines)

	// until there are no more lines to scan
	for lineScanner.Scan() {

		if y > maxY {
			break
		}

		line := lineScanner.Text()
		lastSpaceIndex := -1
		lastSpaceScreen := -1
		size := 1
		r := ' '
		runes := []rune(line)

		// if line is empty, advance y, reset x, keep scanning
		if len(line) == 0 {
			y++
			x = minX
			continue
		}

		// iterate through line of text
		for i := 0; i < len(runes); i++ {

			if y > maxY {
				break
			}

			r = runes[i]

			size = 1

			if utf8.RuneLen(r) > 1 {
				size = 2
			}

			// if rune is a space, then keep track of its index
			if r == ' ' {
				lastSpaceIndex = i
				lastSpaceScreen = x
			}

			// if there is room to add the current rune
			if x+size <= maxX {
				layer.Grid[x][y].Rune = r
				x += size
				continue
			}

			// if there is no more room to add runes from line, advance y and reset x
			y++
			x = minX

			// if there were no spaces this so far this line, rewind i by 1 so we redraw that character next line
			// otherwise if this character is a space, ignore it and start the next line from the next character
			// otherwise if this character wasn't a space, and there was a space this line, blank out everything
			// from the last space to the end of the line, then keep drawing on the next line from the position of
			// of the last space.
			if lastSpaceIndex == -1 {
				i--
			} else if r != ' ' {
				Paint(layer, lastSpaceScreen, y, x, y, ' ', style)
				i = lastSpaceIndex
			}
		}

		// advance y position after reading a complete line
		y++
	}
}
