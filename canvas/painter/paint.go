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
	"unicode/utf8"

	"github.com/apsdsm/canvas"
	"github.com/gdamore/tcell"
)

// Paint covers the rectangle region specified by the two coord pairs with the supplied rune.
//
// Because runes can be either 1 or 2 spaces wide when painted on the screen, the method first
// checks to see how many cells should be advanced in the x axis each call to SetContent before
// it starts painting. If it turns out that it's impossible to paint the requested rune in the
// remaining area, a blank space is painted instead.
func Paint(layer *canvas.Layer, startX, startY, endX, endY int, char rune, style tcell.Style) {

	size := 1

	if utf8.RuneLen(char) > 1 {
		size = 2
	}

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x += size {

			if x == layer.MaxX && size > 1 {
				layer.SetRune(x, y, ' ')
				break
			}

			layer.SetRune(x, y, char)
		}
	}
}
