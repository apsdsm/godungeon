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

package file

import (
	"encoding/json"
	"io/ioutil"

	"github.com/apsdsm/godungeon/file/json_format"
	"github.com/apsdsm/godungeon/game"
)

// LoadMap will load a map file into a map object
func LoadMap(path string) *game.Dungeon {

	in := json_format.Dungeon{}
	infile, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(infile, &in)

	m := game.NewDungeon(in.Width, in.Height)

	// copy tile data
	for x := 0; x < len(in.Tiles); x++ {
		for y := 0; y < len(in.Tiles[x]); y++ {
			m.Tiles[x][y].Rune = in.Tiles[x][y].Rune
			m.Tiles[x][y].Walkable = in.Tiles[x][y].Walkable
			m.Tiles[x][y].Position = game.Position{x, y}
			// @todo spawn actors
		}
	}

	// jump map to calculate neighbor positions:
	// 8 1 2
	// 7   3
	// 6 5 4
	jump := []game.Position{
		{0, -1},  // N
		{1, -1},  // NE
		{1, 0},   // E
		{1, 1},   // SE
		{0, 1},   // S
		{-1, 1},  // SW
		{-1, 0},  // W
		{-1, -1}, // NW
	}

	// assign neighbors
	for x := range m.Tiles {
		for y := range m.Tiles[x] {
			for n := 0; n < 8; n++ {
				nPos := game.Position{x + jump[n].X, y + jump[n].Y}

				if !nPos.OutOfBounds(m.Width, m.Height) {
					m.Tiles[x][y].Neighbors[n] = m.At(nPos.X, nPos.Y)
				}
			}
		}
	}

	// @todo count other entities
	numActors := 1

	// @todo allocate space for all entities, not just player
	m.Actors = make([]game.Actor, numActors, numActors)

	// setup player
	m.Actors[0] = game.Actor{
		Tile:       (m.At(in.StartPosition.X, in.StartPosition.Y)),
		Name:       "player",
		Appearance: 'x',
	}

	// @todo setup other actors

	return &m
}
