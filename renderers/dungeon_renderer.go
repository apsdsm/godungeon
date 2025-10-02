//    Copyright 2017 Nick del Pozo
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

package renderers

import (
	"github.com/apsdsm/godungeon/canvas"
	"github.com/apsdsm/godungeon/game"
	"github.com/gdamore/tcell"
)

// A DungeonRenderer renders the map portion of a actor to a layer
type DungeonRenderer struct {
	player  *game.Actor
	dungeon *game.Dungeon
	layer   *canvas.Layer
}

// DungeonRendererConfig has setup data for new method
type DungeonRendererConfig struct {
	Player  *game.Actor
	Dungeon *game.Dungeon
	Layer   *canvas.Layer
}

// NewDungeonRenderer creates and returns a pointer to a new MapDislayer object
func NewDungeonRenderer(config DungeonRendererConfig) *DungeonRenderer {
	d := DungeonRenderer{}
	d.dungeon = config.Dungeon
	d.layer = config.Layer
	d.player = config.Player

	return &d
}

// Render the map to the screen
func (d *DungeonRenderer) Render() {
	d.layer.Clear()

	// assume the camera is at 0,0, and has a size equal to the layer size.

	camX, camY := 0, 0
	camW, camH := d.layer.Width, d.layer.Height

	// now, which part of the dungeon can we actually see?

	// ensure camera is in bounds
	mapXs := camX
	// mapXe := camX + camW

	if mapXs+(camW-1) >= d.dungeon.Width {
		mapXs = d.dungeon.Width - camW
		// mapXe = d.dungeon.Width - 1
	}

	// start go    from from start
	// stop  stop  to   stop until

	mapYs := camY
	// mapYe := camY + camH

	if mapYs+(camH-1) >= d.dungeon.Height {
		mapYs = d.dungeon.Height - camH
		// mapYe = d.dungeon.Height - 1
	}

	for x := 0; x < d.layer.Width; x++ {
		for y := 0; y < d.layer.Height; y++ {

			// what tile at we looking at?
			mx := mapXs + x
			my := mapYs + y

			d.layer.Grid[x][y].Rune = d.dungeon.Tiles[mx][my].Rune
			d.layer.Grid[x][y].Style = tcell.StyleDefault.Background(tcell.ColorBlack)
		}
	}

	// // clear visibility of tiles
	// for x := 0; x < len(d.dungeon.Tiles); x++ {
	// 	for y := 0; y < len(d.dungeon.Tiles[x]); y++ {
	// 		d.dungeon.Tiles[x][y].Visible = false
	// 	}
	// }

	// math2d.FindVisibleTiles2(d.player.Tile, d.dungeon.Tiles)

	// visbStyle := tcell.StyleDefault.Foreground(game.White).Background(tcell.ColorGrey)
	// seenStyle := tcell.StyleDefault.Foreground(game.Grey).Background(tcell.ColorRed)

	// layerX := 0
	// layerY := 0
	// tileX := mapXs
	// tileY := mapYs

	// for layerX := 0, mapX := mapXs; x < d.layer.Width; layerX++, mapX++ {

	// 	for y := 0, dy := mapYs; y < d.layer.Height; y++, dy++ {
	// 		if d.dungeon.Tiles[dx][dy].Visible {
	// 			d.layer.Grid[layerX][y].Rune = d.dungeon.Tiles[dx][dy].Rune
	// 			d.layer.Grid[layerX][y].Style = visbStyle
	// 			d.dungeon.Tiles[dx][dy].Seen = true
	// 		} else if d.dungeon.Tiles[dx][dy].Seen {
	// 			d.layer.Grid[layerX][y].Rune = d.dungeon.Tiles[dx][dy].Rune
	// 			d.layer.Grid[layerX][y].Style = seenStyle
	// 		}
	// 	}
	// }
}
