package controllers

import (
	"os"

	"github.com/gdamore/tcell"
)

type GameController struct {
	screen tcell.Screen
}

func NewGameController(screen tcell.Screen) GameController {
	c := GameController{
		screen: screen,
	}
	return c
}

// Quit will close the screen and exit the gam
// @todo test this - there are ways of testing os.Exit but they all seem pretty hacky?
func (c *GameController) Quit() {
	c.screen.Fini()
	os.Exit(0)
}
