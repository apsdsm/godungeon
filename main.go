package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/canvas"
	"github.com/apsdsm/canvas/painter"

	"github.com/gdamore/tcell"

	"github.com/apsdsm/godungeon/controllers"
	"github.com/apsdsm/godungeon/debug"
	"github.com/apsdsm/godungeon/file"
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/input"
	"github.com/apsdsm/godungeon/renderers"
	"github.com/apsdsm/godungeon/updaters"
)

var (
	screen        tcell.Screen
	width, height int
	gameCanvas    canvas.Canvas
	mapLayer      canvas.Layer
	entityLayer   canvas.Layer
	consoleLayer  canvas.Layer
)

func main() {
	// initialize the screen
	screen, width, height = createAndInitScreen()

	// create game canvas
	gameCanvas = canvas.NewCanvas(screen)

	// create game layers
	mapLayer = canvas.NewLayer(width, height, 0, 0)
	entityLayer = canvas.NewLayer(width, height, 0, 0)
	consoleLayer = canvas.NewLayer(width, 10, 0, height-11)

	// add layers to canvas
	gameCanvas.AddLayer(&mapLayer)
	gameCanvas.AddLayer(&entityLayer)
	gameCanvas.AddLayer(&consoleLayer)

	// set up input
	inputHandler := input.NewHandler(screen)

	// set up objects
	dungeon := file.LoadMap("fixtures/maps/simple.json")

	// set up renderers
	mapRenderer := renderers.NewDungeonRenderer(renderers.DungeonRendererConfig{
		Dungeon: dungeon,
		Layer:   &mapLayer,
		Player:  &dungeon.Actors[0],
	})

	entityRenderer := renderers.NewActorRenderer(
		&dungeon.Actors,
		&entityLayer,
	)

	// set up controllers
	actorController := controllers.NewActorController(controllers.ActorControllerConfig{})

	// set up updaters
	player := updaters.NewPlayer(&dungeon.Actors[0], &inputHandler, &actorController)

	mobs := updaters.NewMobAi(updaters.MobAiConfig{
		Player: &dungeon.Actors[0],
		Mobs:   &dungeon.Actors,
	})

	// bind player movement <- should be in config object, loaded from config file
	player.BindMovement(input.NewKey(input.KeyUp, 0), game.N)
	player.BindMovement(input.NewKey(input.KeyRight, 0), game.E)
	player.BindMovement(input.NewKey(input.KeyDown, 0), game.S)
	player.BindMovement(input.NewKey(input.KeyLeft, 0), game.W)

	// initial render of content to layers
	mapRenderer.Render()
	entityRenderer.Render()
	gameCanvas.Draw()

	// main game loop
	for {
		// update input
		inputHandler.Update()

		// quit if user presses 'q' <- temporary code until a main menu system is in place
		if inputHandler.HasKeyEvent(input.NewKey(input.KeyRune, 'q')) {
			exitGame()
		}

		// update updaters <- should be triggering these from a loop
		player.Update()
		mobs.Update()

		// lear layer <- should only do this if dirty
		entityLayer.Clear()

		// update renderers <- should only do this if dirty
		entityRenderer.Render()
		mapRenderer.Render()

		// render the cosole <- should only do this if dirty
		renderConsole()
		gameCanvas.Draw()
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

// cleanly exit the game
func exitGame() {
	screen.Fini()
	os.Exit(1)
}

func renderConsole() {
	consoleLayer.Clear()
	log := debug.Tail(10)

	for i := range log {
		painter.DrawText(&consoleLayer, 0, i, log[i], tcell.StyleDefault.Foreground(tcell.Color104))
	}
}
