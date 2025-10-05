//    Copyright 2017 nick
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

// DrawBox will draw a box on the layer
func DrawBox(layer *canvas.Layer, startX, startY, endX, endY int, style tcell.Style) {

	// fmt.Println(startX)
	// fmt.Println(startY)
	// fmt.Println(endX)
	// fmt.Println(endY)

	// draw corners using set rune
	layer.SetRune(startX, startY, '+')
	layer.SetRune(endX, startY, '+')
	layer.SetRune(endX, endY, '+')
	layer.SetRune(startX, endY, '+')

	// draw top and bottom lines
	for i := startX + 1; i < endX; i++ {
		layer.SetRune(i, startY, '-')
		layer.SetRune(i, endY, '-')
	}

	// draw left and right lines
	for i := startY + 1; i < endY; i++ {
		layer.SetRune(startX, i, '|')
		layer.SetRune(endX, i, '|')
	}
}
