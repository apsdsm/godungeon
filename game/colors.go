package game

import "github.com/gdamore/tcell"

// Color is a generic type
type Color = tcell.Color

// This table defines local lookups that translate to tcell colors.
const (
	White    = tcell.ColorWhite
	Grey     = tcell.ColorGrey
	DarkGrey = tcell.ColorDarkGrey
	Green    = tcell.ColorGreen
	Red      = tcell.ColorRed
	Orange   = tcell.ColorOrange
)
