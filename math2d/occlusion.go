package math2d

import (
	"github.com/apsdsm/godungeon/game"
)

// FindVisibleTiles finds all the tiles which are visible from a given starting point
func FindVisibleTiles(start *game.Tile, tiles [][]game.Tile) {

	// make a map of which tiles are in checkable range
	maxDist := 10

	potentials := game.TilesInRange(start, tiles, maxDist)

	// set up some constant values that are used to quickly calc the verts
	// surrouding each tile.
	//
	// ~~set the points slightly inset forom the tile.~~
	//
	// now, this is a hack, but get this - if you set the four
	// points to check to be exactly the four corners of the tile, then
	// you run into problems with false positives - where you are actually
	// checking and failing the sight line based on the visibility of
	// adjacent tiles that do not nominally come into the picture. By
	// setting the sightlines to be slightly inside the tile, you
	// never hit those troublesome tiles...
	//
	// but, it's still a hack, and won't work for corners.
	// what you really need to do is check how far along the line the intersection
	// is. if line A intersects line B at its very farthest point, then
	// we know that we can probably not worry about that tile.
	points := []game.Vec2{
		{-0.49, -0.49},
		{0.49, -0.49},
		{0.49, 0.49},
		{-0.49, 0.49},
	}

	// generate the vector for the start tile
	startVec := game.Vec2{float64(start.Position.X), float64(start.Position.Y)}

	// for each potentially visible tile, check to see if we can draw a line
	// from the start tile to that line where each tile intersected by that
	// line is walkable. If even a single intersecting tile is not walkable,
	// then that line is deemed not visible. However, if a clean line can be
	// drawn to the potential, that tile is deemed visible.
	for _, p := range potentials {

		// calculate the points around the potential
		tileVecs := []game.Vec2{
			{float64(p.Position.X) + points[0].X, float64(p.Position.Y) + points[0].Y},
			{float64(p.Position.X) + points[1].X, float64(p.Position.Y) + points[1].Y},
			{float64(p.Position.X) + points[2].X, float64(p.Position.Y) + points[2].Y},
			{float64(p.Position.X) + points[3].X, float64(p.Position.Y) + points[3].Y},
		}

		// the lines that lead from the start tile to the four corners of this tile
		linesToTile := []game.Line{
			{startVec, tileVecs[0]},
			{startVec, tileVecs[1]},
			{startVec, tileVecs[2]},
			{startVec, tileVecs[3]},
		}

		var visible bool

		for _, line := range linesToTile {
			visible = true

			for _, p2 := range potentials {
				if p2 == p {
					continue
				}

				if game.LineIntersectsTile(line, p2) {
					if !p2.Walkable {
						visible = false
						break
					}
				}
			}

			if visible {
				break
			}
		}

		if visible {
			p.Visible = true
		}
	}
}
