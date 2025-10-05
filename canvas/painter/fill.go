package painter

import (
	"unicode/utf8"

	"github.com/apsdsm/canvas"
	"github.com/gdamore/tcell"
)

// Fill covers the entire layer with the specified rune.
//
// Because runes can be either 1 or 2 spaces wide when painted on the screen, the method first
// checks to see how many cells should be advanced in the x axis each call to SetContent before
// it starts painting. If it turns out that it's impossible to paint the requested rune in the
// remaining area, a blank space is painted instead.
func Fill(layer *canvas.Layer, char rune, style tcell.Style) {

	size := 1

	if utf8.RuneLen(char) > 1 {
		size = 2
	}

	for x := 0; x < layer.Width; x += size {
		for y := 0; y < layer.Height; y++ {

			if x == layer.MaxX && size > 1 {
				layer.SetRune(x, y, ' ')
				break
			}

			layer.SetRune(x, y, char)

			if size == 2 {
				layer.SetRune(x+1, y, ' ')
			}
		}
	}
}
