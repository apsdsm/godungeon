package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/canvas"

	"github.com/gdamore/tcell"

	"github.com/apsdsm/godungeon/controllers"
	"github.com/apsdsm/godungeon/file"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
	"github.com/apsdsm/godungeon/renderers/actor_renderer"
	"github.com/apsdsm/godungeon/renderers/dungeon_renderer"
	"github.com/apsdsm/godungeon/updaters"
)

func main() {
	// initialize the screen
	screen, width, height := createAndInitScreen()

	// create game canvas and layers
	gameCanvas := canvas.NewCanvas(screen)
	mapLayer := canvas.NewLayer(width, height, 0, 0)
	entityLayer := canvas.NewLayer(width, height, 0, 0)
	gameCanvas.AddLayer(&mapLayer)
	gameCanvas.AddLayer(&entityLayer)

	// load map
	dungeon := file.LoadMap("fixtures/maps/simple.json")

	// setup an input handler
	inputHandler := input.NewHandler(screen)

	// set up map renderer
	mapRenderer := dungeon_renderer.New(dungeon, &mapLayer)

	// set up entity renderer
	entityRenderer := actor_renderer.New(&dungeon.Actors, &entityLayer)

	// set up a player
	player := updaters.NewPlayer(&dungeon.Actors[0], &inputHandler)
	player.BindMovement(input.NewKey(input.KeyUp, 0), game.N)
	player.BindMovement(input.NewKey(input.KeyRight, 0), game.E)
	player.BindMovement(input.NewKey(input.KeyDown, 0), game.S)
	player.BindMovement(input.NewKey(input.KeyLeft, 0), game.W)

	// initial render of content to layers
	mapRenderer.Render()
	entityRenderer.Render()
	gameCanvas.Draw()

	// make controllers used directly in scene (this should be moved to an updated)
	game := controllers.NewGameController(screen)

	// main game loop <- move this logic into a specific scene object, rather than the main loop
	for {

		// update input (I wish this were in an actual game loop object)
		inputHandler.Update()

		// quit if user presses 'q' <- temporary code until a main menu system is in place
		if inputHandler.HasKeyEvent(input.NewKey(input.KeyRune, 'q')) {
			game.Quit()
		}

		// updaters <- should be triggering these from a loop
		player.Update()

		//@todo only render if dirty (move this to actor object)
		//if dirty {
		entityLayer.Clear()
		entityRenderer.Render()
		gameCanvas.Draw()
		//}
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
