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

package game

// A Dungeon contains a grid of tiles
type Dungeon struct {
	Width  int
	Height int
	Link   string
	Tiles  [][]Tile
	Actors []Actor
	ActorPrototypes []Actor
}

// NewDungeon generates a new dungeon initialized to the specified size
func NewDungeon(width, height int) Dungeon {
	d := Dungeon{}

	d.Width = width
	d.Height = height
	d.Tiles = make([][]Tile, width)

	for i := range d.Tiles {
		d.Tiles[i] = make([]Tile, height)
	}

	return d
}

// At returns the tile at the given coords
func (d *Dungeon) At(x, y int) *Tile {
	return &d.Tiles[x][y]
}

// InBounds returns true if the position is inside the dungeon
func (d *Dungeon) InBounds(x, y int) bool {
	return 0 <= x && x < d.Width && 0 <= y && y < d.Height
}
