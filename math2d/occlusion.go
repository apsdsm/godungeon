package math2d

import (
	"github.com/apsdsm/godungeon/game"
)

func FindVisibleTilesRad(start *game.Tile, tiles [][]game.Tile) {
	// maxDist := 10
	// potentials := game.TilesInRange(start, tiles, maxDist)

	// rays := 100

	// incr := 360.0 / float64(rays)

	// angle := float64(0)

	// rayLines := make([]game.Line, rays)

	// up :=

	// for i := 0; i < rays; i++ {
	// 	angleRads := (math.Pi / 180) * angle

	// 	angle += incr
	// }

}

// FindVisibleTiles finds all the tiles which are visible from a given starting point
func FindVisibleTiles(start *game.Tile, tiles [][]game.Tile) {

	// make a map of which tiles are in checkable range
	maxDist := 10

	potentials := game.TilesInRange(start, tiles, maxDist)

	// set up some constant values that are used to quickly calc the verts
	// surrouding each tile.
	//
	// set 8 points - 4 slightly inside the tile, and 4 slightly outside.
	//
	// now, this is a hack, but get this - if you set the four
	// points to check to be exactly the four corners of the tile, then
	// you run into problems with false positives - where you are actually
	// checking and failing the sight line based on the visibility of
	// adjacent tiles that do not nominally come into the picture. By
	// setting the sightlines to be slightly inside the tile, you
	// never hit those troublesome tiles, and by setting an additional
	// four points just *outside* the tile you additionally cover the
	// literal corner case where one tile is obscured by two corners.
	//
	// but, it's still a hack, and it means you're checking each points on
	// each tile, but these checks are dirt cheap, so we can get away with
	// being a little permissive here. If it really gets performance intensive
	// then later on it could be possible to create baked in visibility
	// maps for each level.
	points := []game.Vec2{
		{-0.49, -0.49},
		{0.49, -0.49},
		{0.49, 0.49},
		{-0.49, 0.49},
		{-0.51, -0.51},
		{0.51, -0.51},
		{0.51, 0.51},
		{-0.51, 0.51},
	}

	// vector of start tile position
	startVec := start.Center()

	// for each potentially visible tile, check to see if we can draw a line
	// from the start tile to that line where each tile intersected by that
	// line is walkable. If even a single intersecting tile is not walkable,
	// then that line is deemed not visible. However, if a clean line can be
	// drawn to the potential, that tile is deemed visible.
	for _, p := range potentials {

		var tileVecs []game.Vec2
		var linesToTile []game.Line

		// if p.Walkable {
		// 	tileVecs = []game.Vec2{
		// 		p.Center(),
		// 	}
		// } else {
		// 	a := p.Coords()
		// 	tileVecs = a[:]
		// }

		// calculate the points around the potential
		tileVecs = []game.Vec2{
			{float64(p.Position.X) + points[0].X, float64(p.Position.Y) + points[0].Y},
			{float64(p.Position.X) + points[1].X, float64(p.Position.Y) + points[1].Y},
			{float64(p.Position.X) + points[2].X, float64(p.Position.Y) + points[2].Y},
			{float64(p.Position.X) + points[3].X, float64(p.Position.Y) + points[3].Y},
			{float64(p.Position.X) + points[4].X, float64(p.Position.Y) + points[4].Y},
			{float64(p.Position.X) + points[5].X, float64(p.Position.Y) + points[5].Y},
			{float64(p.Position.X) + points[6].X, float64(p.Position.Y) + points[6].Y},
			{float64(p.Position.X) + points[7].X, float64(p.Position.Y) + points[7].Y},
		}

		// tileCoords := p.Coords()
		// tileVecs = tileCoords[:]

		//make a line from each check vector to the center of the start tile
		for v := range tileVecs {
			linesToTile = append(linesToTile, game.Line{startVec, tileVecs[v]})
		}

		// // // the lines that lead from the start tile to the four corners of this tile
		// linesToTile = []game.Line{
		// 	{startVec, tileVecs[0]},
		// 	{startVec, tileVecs[1]},
		// 	{startVec, tileVecs[2]},
		// 	{startVec, tileVecs[3]},
		// 	{startVec, tileVecs[4]},
		// 	{startVec, tileVecs[5]},
		// 	{startVec, tileVecs[6]},
		// 	{startVec, tileVecs[7]},
		// }

		var visible bool

		for _, line := range linesToTile {
			visible = true

			for _, p2 := range potentials {
				if p2 == p {
					continue
				}

				if game.LineIntersectsTileTol(line, p2, 0) {
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
