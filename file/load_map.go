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

	"strconv"

	"github.com/apsdsm/godungeon/debug"
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

	// this is a temporary memory allocation. This function should look through the
	// map and make an acurate allocation based on how many enemies need to be spawned.
	m.Actors = make([]game.Actor, 1, 50)

	// setup player <- should pull this data from the json file
	m.Actors[0] = game.Actor{
		Tile:     m.At(in.StartPosition.X, in.StartPosition.Y),
		Name:     "player",
		Rune:     'x',
		IsPlayer: true,
		Hp:       20,
		Attack: game.Attack{
			MinDamage:   5,
			MaxDamage:   10,
			ChanceToHit: 100,
		},
	}

	// setup tiles
	for x := 0; x < len(in.Tiles); x++ {
		for y := 0; y < len(in.Tiles[x]); y++ {
			// tile data
			m.Tiles[x][y].Rune = in.Tiles[x][y].Rune
			m.Tiles[x][y].Walkable = in.Tiles[x][y].Walkable
			m.Tiles[x][y].Position = game.NewPosition(x, y)

			// spawn actor in tile
			if in.Tiles[x][y].Spawn != "" {
				a := makeActor(in.Tiles[x][y].Spawn, &in.Mobs)
				a.Tile = m.At(x, y)
				m.Actors = append(m.Actors, a)
				m.Tiles[x][y].Occupant = &m.Actors[len(m.Actors)-1]
			}
		}
	}

	// jump map to calculate neighbor positions:
	// 8 1 2
	// 7   3
	// 6 5 4
	jump := []game.Position{
		{X: 0, Y: -1},  // N
		{X: 1, Y: -1},  // NE
		{X: 1, Y: 0},   // E
		{X: 1, Y: 1},   // SE
		{X: 0, Y: 1},   // S
		{X: -1, Y: 1},  // SW
		{X: -1, Y: 0},  // W
		{X: -1, Y: -1}, // NW
	}

	// assign neighbors
	for x := range m.Tiles {
		for y := range m.Tiles[x] {
			for n := 0; n < 8; n++ {
				nPos := game.Position{
					X: x + jump[n].X,
					Y: y + jump[n].Y,
				}

				if !nPos.OutOfBounds(m.Width, m.Height) {
					m.Tiles[x][y].Neighbors[n] = m.At(nPos.X, nPos.Y)
				}
			}
		}
	}

	return &m
}

// resolvePrototypes takes a mob object, then recurses up the prototype chain, filling in
// values that are currently blank as it finds them. This ensures that the values closest
// to the base object are used, while allowing for multiple levels of inheritance. If a
// mob requires a prototype that doesn't exist, the method will panic.
func resolvePrototypes(mob json_format.Mob, mobs *[]json_format.Mob) json_format.Mob {
	if mob.Prot == "" {
		return mob
	}

	prot := json_format.Mob{}

	for _, p := range *mobs {
		if mob.Prot == p.Link {
			if p.Prot != "" {
				prot = resolvePrototypes(prot, mobs)
			} else {
				prot = p
			}

			if mob.Rune == "" {
				mob.Rune = prot.Rune
			}

			if mob.Name == "" {
				mob.Name = prot.Name
			}

			if mob.Link == "" {
				mob.Link = prot.Link
			}

			if mob.Hp == "" {
				mob.Hp = prot.Hp
			}

			if mob.Mp == "" {
				mob.Mp = prot.Mp
			}

			return mob
		}
	}

	panic("was not able to resolve prototype: " + mob.Prot)
}

// makeActor will search for a mob with the given link, and return an Actor initialized to
// those values. If the link is not contained in the array of mobs, the method will panic.
func makeActor(link string, mobs *[]json_format.Mob) game.Actor {
	actor := game.Actor{}

	for _, m := range *mobs {
		if m.Link == link {
			m = resolvePrototypes(m, mobs)
			actor.Name = m.Name
			actor.Link = m.Link
			actor.Rune = parseRune(m.Rune)
			actor.Hp = parseInt(m.Hp)
			actor.MaxHp = actor.Hp
			actor.Mp = parseInt(m.Mp)
			actor.MaxMp = actor.Mp

			return actor
		}
	}

	panic("was not able to build the actor: " + link)
}

// parseRune converts a string to a rune. It uses the first character of the string. If
// conversion is impossible the method will panic.
func parseRune(s string) rune {
	if s != "" {
		return rune(s[0])
	}
	panic("could not parse rune")
}

// parseInt converts a string to an int. If conversation fails the method will panic.
func parseInt(s string) int {
	if s == "" {
		debug.Log("empty string == 0")
		return 0
	}

	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		panic("could not parse int")
	}

	return int(val)
}
