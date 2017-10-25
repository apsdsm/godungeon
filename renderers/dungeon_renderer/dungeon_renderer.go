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

package dungeon_renderer

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/game"
)

// A DungeonRenderer renders the map portion of a actor to a layer
type DungeonRenderer struct {
	dungeon *game.Dungeon
	layer   *canvas.Layer
}

// New creates and returns a pointer to a new MapDislayer object
func New(dungeon *game.Dungeon, layer *canvas.Layer) *DungeonRenderer {
	d := DungeonRenderer{}
	d.dungeon = dungeon
	d.layer = layer
	return &d
}

// Render the map to the screen
func (d *DungeonRenderer) Render() {
	for x := 0; x < len(d.dungeon.Tiles); x++ {
		for y := 0; y < len(d.dungeon.Tiles[x]); y++ {
			d.layer.Grid[x][y].Rune = d.dungeon.Tiles[x][y].Rune
		}
	}
}
