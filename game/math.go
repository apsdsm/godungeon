package game

import (
	"math"
)

// TDist returns the number of steps between two tiles (assuming a diagonal step is 2)
func TDist(from, to *Tile) int {
	xdis := math.Abs(float64(from.Position.X) - float64(to.Position.X))
	ydis := math.Abs(float64(from.Position.Y) - float64(to.Position.Y))

	return int(xdis + ydis)
}

type Vec2 struct {
	X float64
	Y float64
}

func Vec2Sub(a, b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

func Vec2Add(a, b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func PosAddVec2(a Position, b Vec2) Vec2 {
	return Vec2{float64(a.X) + b.X, float64(a.Y) + b.Y}
}

func Vec2MulF(a Vec2, b float64) Vec2 {
	return Vec2{a.X * b, a.Y * b}
}

func Vec2Mag(a Vec2) float64 {
	return a.X*a.X + a.Y*a.Y
}

func Vec2Len(a Vec2) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

func Vec2Nor(a Vec2) Vec2 {
	len := Vec2Len(a)
	return Vec2{a.X / len, a.Y / len}
}

func Vec2Dot(a, b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

type pftile struct {
	tile   *Tile
	dist   int
	parent *Tile
	open   bool
}

func TDeg(from, to *Tile) float64 {
	b := Vec2{float64(to.Position.X), float64(to.Position.Y)}
	a := Vec2{float64(from.Position.X), float64(from.Position.Y)}
	ba := Vec2Sub(b, a)
	norm := Vec2Nor(ba)
	deg := math.Atan2(norm.X, norm.Y)*180/math.Pi + 180.0

	return deg
}

func TVis(from, to *Tile) bool {

	return true
}

func LinesIntersect(l1, l2 Line) bool {
	x1 := l1.A.X
	y1 := l1.A.Y
	x2 := l1.B.X
	y2 := l1.B.Y
	x3 := l2.A.X
	y3 := l2.A.Y
	x4 := l2.B.X
	y4 := l2.B.Y

	denom := ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))
	uA := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denom
	uB := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / denom

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true
	}
	return false
}

func LineIntersectsTile(line Line, tile *Tile) bool {

	// when vecs added to tile pos will get the 4 tile points
	points := []Vec2{
		{-0.5, -0.5},
		{0.5, -0.5},
		{0.5, 0.5},
		{-0.5, 0.5},
	}

	// indices of the tile points that make the tile lines
	lines := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
	}

	tileVecs := []Vec2{
		{float64(tile.Position.X) + points[0].X, float64(tile.Position.Y) + points[0].Y},
		{float64(tile.Position.X) + points[1].X, float64(tile.Position.Y) + points[1].Y},
		{float64(tile.Position.X) + points[2].X, float64(tile.Position.Y) + points[2].Y},
		{float64(tile.Position.X) + points[3].X, float64(tile.Position.Y) + points[3].Y},
	}

	// calculate the lines around this tile
	tileLines := []Line{
		{tileVecs[lines[0][0]], tileVecs[lines[0][1]]},
		{tileVecs[lines[1][0]], tileVecs[lines[1][1]]},
		{tileVecs[lines[2][0]], tileVecs[lines[2][1]]},
		{tileVecs[lines[3][0]], tileVecs[lines[3][1]]},
	}

	for _, tileLine := range tileLines {
		if LinesIntersect(line, tileLine) {
			return true
		}
	}

	return false
}

func TilesInRange(start *Tile, tiles [][]Tile, r int) []*Tile {
	estSize := len(tiles) * len(tiles[0])
	inRange := make([]*Tile, 0, estSize)

	for x := 0; x < len(tiles); x++ {
		for y := 0; y < len(tiles[x]); y++ {
			if TDist(start, &tiles[x][y]) <= r {
				inRange = append(inRange, &tiles[x][y])
			}
		}
	}

	return inRange
}
