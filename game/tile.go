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

var tileOffsets = [...]Vec2{
	{-0.5, -0.5},
	{0.5, 0.5},
	{0.5, -0.5},
	{-0.5, 0.5},
}

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

// Center returns the vec2 at the center of this tile
func (t *Tile) Center() Vec2 {
	return Vec2{
		X: float64(t.Position.X),
		Y: float64(t.Position.Y),
	}
}

// Coords returns the coordinates surrounding this tile
func (t *Tile) Coords() [4]Vec2 {
	vec := t.Center()
	var coords [4]Vec2
	coords[0] = Vec2{vec.X + tileOffsets[0].X, vec.Y + tileOffsets[0].Y}
	coords[1] = Vec2{vec.X + tileOffsets[1].X, vec.Y + tileOffsets[1].Y}
	coords[2] = Vec2{vec.X + tileOffsets[2].X, vec.Y + tileOffsets[2].Y}
	coords[3] = Vec2{vec.X + tileOffsets[3].X, vec.Y + tileOffsets[3].Y}
	return coords
}
