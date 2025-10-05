package canvas

import "github.com/gdamore/tcell"

// A Cell is a single point in a layer (a single char space on the screen)
type Cell struct {
	Rune  rune
	Style tcell.Style
}

func (c *Cell) Set (Rune rune, Style tcell.Style) {
	c.Rune = Rune
	c.Style = Style
}
