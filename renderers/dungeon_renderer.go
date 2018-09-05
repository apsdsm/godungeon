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
	"github.com/apsdsm/canvas"
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

	// test player pos to origin
	//game.TVis(d.player.Tile, &d.dungeon.Tiles[2][1])

	for x := 0; x < len(d.dungeon.Tiles); x++ {
		for y := 0; y < len(d.dungeon.Tiles[x]); y++ {

			dist := game.TDist(d.player.Tile, &d.dungeon.Tiles[x][y])

			style := tcell.StyleDefault

			if dist <= d.player.Sight && game.TVis(d.player.Tile, &d.dungeon.Tiles[x][y]) {
				d.dungeon.Tiles[x][y].Seen = true
				d.layer.Grid[x][y].Rune = d.dungeon.Tiles[x][y].Rune
				d.layer.Grid[x][y].Style = style.Foreground(game.White).Background(game.Orange)

				if d.dungeon.Tiles[x][y].Walkable {
					d.layer.Grid[x][y].Rune = ' '
				}

			} else if d.dungeon.Tiles[x][y].Seen {
				d.layer.Grid[x][y].Rune = d.dungeon.Tiles[x][y].Rune
				d.layer.Grid[x][y].Style = style.Foreground(game.Grey).Background(game.DarkGrey)

				if d.dungeon.Tiles[x][y].Walkable {
					d.layer.Grid[x][y].Rune = '·'
				}
				
			} else {				
				d.layer.Grid[x][y].Rune = d.dungeon.Tiles[x][y].Rune
				d.layer.Grid[x][y].Style = style.Foreground(game.Grey).Background(game.DarkGrey)
				d.layer.Grid[x][y].Rune = '·'
			}
		}
	}
}
