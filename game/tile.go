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

// A Tile is a single tile in a map
type Tile struct {
	Rune       rune
	Walkable   bool
	Spawn      string
	Occupant   *Actor
	Position   Position
	Neighbors  [8]*Tile
	Seen       bool
	Visible    bool
	Brightness float32
}

// Neighbor returns the neight of this tile in the given direction
func (t *Tile) Neighbor(direction Direction) *Tile {
	return t.Neighbors[direction]
}
