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
	d.makePotentialList(d.player.Tile)

	style := tcell.StyleDefault
	style.Foreground(game.White)

	for x := 0; x < len(d.dungeon.Tiles); x++ {
		for y := 0; y < len(d.dungeon.Tiles[x]); y++ {

			if !d.dungeon.Tiles[x][y].Visible {
				continue
			}

			d.layer.Grid[x][y].Style = style

			if d.dungeon.Tiles[x][y].Walkable {
				d.layer.Grid[x][y].Rune = ' '
			} else {
				d.layer.Grid[x][y].Rune = d.dungeon.Tiles[x][y].Rune
			}
		}
	}
}

func (d *DungeonRenderer) makePotentialList(start *game.Tile) {

	// get a map of which tiles are visible
	potentials := make(map[*game.Tile]bool)
	potentials[start] = false
	seeking := true
	maxDist := 10

	// get a list of which tiles are in visible range of the player. this
	// creates a subset of the tiles which need to be checked in visibility
	// calculations
	for seeking == true {
		seeking = false
		for p, checked := range potentials {
			if !checked {
				for _, n := range p.Neighbors {
					if n == nil {
						continue
					}

					if _, exists := potentials[n]; !exists {
						if game.TDist(start, n) < maxDist {
							potentials[n] = false
							seeking = true
						}
					}
				}
				potentials[p] = true
			}
		}
	}

	// set up some constant values that are used to quickly calc the verts
	// surrouding each tile.

	// when vecs added to tile pos will get the 4 tile points
	points := []game.Vec2{
		{-0.5, -0.5},
		{0.5, -0.5},
		{0.5, 0.5},
		{-0.5, -0.5},
	}

	// indices of the tile points that make the tile lines
	lines := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
	}

	// make a list of which tiles are visible
	visibles := make(map[*game.Tile]bool)
	visibles[start] = true
	startVec := game.Vec2{float64(start.Position.X), float64(start.Position.Y)}

	// for each potential, check if it's visible
	for p := range potentials {
		if _, alreadyChecked := visibles[p]; alreadyChecked {
			continue
		}

		// calculate the points around this tile
		tileVecs := []game.Vec2{
			{float64(p.Position.X) + points[0].X, float64(p.Position.Y) + points[0].Y},
			{float64(p.Position.X) + points[1].X, float64(p.Position.Y) + points[1].Y},
			{float64(p.Position.X) + points[2].X, float64(p.Position.Y) + points[2].Y},
			{float64(p.Position.X) + points[3].X, float64(p.Position.Y) + points[3].Y},
		}

		// for each of the four vectors made above, draw a line from the center of
		// the start tile to that vector. We will check to see which tiles intersect
		// that line. In order for the line to be marked as clear, all tiles intersected
		// by the line must be walkable. If any line is clear, we stop processing and
		// mark the tile as visible.
		var visible bool
		for _, vec := range tileVecs {

			// start assuming the line is visible
			visible = true

			// this is the line we're going to check
			lineToVec := game.Line{startVec, vec}

			// check the line against all potentials (p2). we are searching for
			// anything that would obstruct the view along this line
			for p2 := range potentials {

				// if the target tile is the same as the potential continue
				if p == p2 {
					continue
				}

				// if the target tile was already checked and is visible then it
				// cannot be an obstruction and we can continue
				if visible, checked := visibles[p2]; checked && visible {
					continue
				}

				// if we got this far then we need to check if the line can pass
				// intersects the tile. if it does then we need to check if the
				// tile would obstruct the view or not

				// calculate the points around this tile
				p2Vecs := []game.Vec2{
					{float64(p2.Position.X) + points[0].X, float64(p2.Position.Y) + points[0].Y},
					{float64(p2.Position.X) + points[1].X, float64(p2.Position.Y) + points[1].Y},
					{float64(p2.Position.X) + points[2].X, float64(p2.Position.Y) + points[2].Y},
					{float64(p2.Position.X) + points[3].X, float64(p2.Position.Y) + points[3].Y},
				}

				// calculate the lines around this tile
				p2Lines := []game.Line{
					{p2Vecs[lines[0][0]], p2Vecs[lines[0][1]]},
					{p2Vecs[lines[1][0]], p2Vecs[lines[1][1]]},
					{p2Vecs[lines[2][0]], p2Vecs[lines[2][1]]},
					{p2Vecs[lines[3][0]], p2Vecs[lines[3][1]]},
				}

				// check each line of the square surrounding the tile to see if
				// line intersects rect
				intersects := false

				for _, line := range p2Lines {
					if game.LinesIntersect(lineToVec, line) {
						intersects = true
						break
					}
				}

				// if the line does not intersect this tile then keep processing
				if !intersects {
					continue
				}

				// if the line to target vector intersects this tile
				// and if the tile is not walkable, then the view is obstructed
				// and we can stop checking this line
				if !p2.Walkable {
					visible = false
					break
				}
			}

			// if we checked all the potential tiles for this line and and there were
			// no obstructions then there is at least one clear line to the target tile.
			// we only need one line to be unobstructed for the tile to be visble so
			// we can stop checking.
			if visible {
				break
			}
		}

		visibles[p] = visible
	}

	for x := range d.dungeon.Tiles {
		for y := range d.dungeon.Tiles[x] {
			t := &d.dungeon.Tiles[x][y]

			if visible, defined := visibles[t]; defined {
				t.Visible = visible
			} else {
				t.Visible = false
			}
		}
	}

}
