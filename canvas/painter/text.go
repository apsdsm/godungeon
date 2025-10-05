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

// DrawText draws text to the x y coords, automatically allocating enough space for double width chars.
func DrawText(layer *canvas.Layer, x, y int, text string, style tcell.Style) {

	// for each rune in the string, draw it on the layer
	for _, r := range text {

		if x > layer.MaxX {
			break
		}

		size := 1

		if utf8.RuneLen(r) > 1 {
			size = 2
		}

		if x == layer.MaxX && size > 1 {
			layer.Grid[x][y].Rune = 0
			break
		}

		layer.Grid[x][y].Rune = r

		x++
	}
}
