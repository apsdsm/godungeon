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
	"github.com/apsdsm/canvas"
	"github.com/gdamore/tcell"
)

// DrawHLine draws a horizontal line on the screen.
func DrawHLine(layer *canvas.Layer, x, y, len int, style tcell.Style) {
	for i := 0; i < len; i++ {

		if x >= layer.Width {
			break
		}

		layer.Grid[x][y].Rune = '-'

		x++
	}
}

// DrawVLine draws a vertical line on the screen.
func DrawVLine(layer *canvas.Layer, x, y, len int, style tcell.Style) {
	for i := 0; i < len; i++ {

		if y >= layer.Height {
			break
		}

		layer.Grid[x][y].Rune = '|'

		y++
	}
}
