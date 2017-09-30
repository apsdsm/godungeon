package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/canvas"

	"github.com/gdamore/tcell"

	"github.com/apsdsm/godungeon/render"

	"github.com/apsdsm/godungeon/io"
)

func main() {
	screen, width, height := createAndInitScreen()

	gameCanvas := canvas.NewCanvas(screen)

	mapLayer := canvas.NewLayer(width, height, 0, 0)
	entityLayer := canvas.NewLayer(width, height, 0, 0)

	gameCanvas.AddLayer(mapLayer)
	gameCanvas.AddLayer(entityLayer)

	dungeon := io.LoadMap("fixtures/game/simple.json")

	mapRenderer := render.NewMapRenderer()
	entityRenderer := render.NewEntityRenderer()

	mapRenderer.DrawMap(&dungeon, mapLayer)
	entityRenderer.DrawPlayer(&dungeon.Player, entityLayer)

	gameCanvas.Draw()

	// main game loop
	for {
		event := screen.PollEvent()

		switch e := event.(type) {

		case *tcell.EventKey:
			if e.Key() == tcell.KeyRune && string(e.Rune()) == "q" {
				screen.Fini()
				os.Exit(0)
			}
			if e.Key() == tcell.KeyLeft {
				dungeon.Player.CurrentPosition.X--
				entityRenderer.DrawPlayer(&dungeon.Player, entityLayer)
				gameCanvas.Draw()
			}
			if e.Key() == tcell.KeyRight {
				dungeon.Player.CurrentPosition.X++
				entityRenderer.DrawPlayer(&dungeon.Player, entityLayer)
				gameCanvas.Draw()
			}
			if e.Key() == tcell.KeyUp {
				dungeon.Player.CurrentPosition.Y--
				entityRenderer.DrawPlayer(&dungeon.Player, entityLayer)
				gameCanvas.Draw()
			}
			if e.Key() == tcell.KeyDown {
				dungeon.Player.CurrentPosition.Y++
				entityRenderer.DrawPlayer(&dungeon.Player, entityLayer)
				gameCanvas.Draw()
			}
		}
	}
}

// create and initialize a tcell screen, die if there are errors
func createAndInitScreen() (screen tcell.Screen, width, height int) {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Println("screen create error")
		os.Exit(1)
	}

	err = screen.Init()

	if err != nil {
		fmt.Println("screen init error")
		os.Exit(1)
	}

	width, height = screen.Size()

	return screen, width, height
}
