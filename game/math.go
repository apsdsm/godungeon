package game

import (
	"math"
)

// when vecs added to tile pos will get the 4 tile points
var tilePoints = []Vec2{
	{-0.5, -0.5},
	{0.5, -0.5},
	{0.5, 0.5},
	{-0.5, 0.5},
}

var tileLines = [][]int{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 0},
}

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

func LinesIntersectTol(l1, l2 Line, tolerance float64) bool {
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

	t0 := 0 + tolerance
	t1 := 1 - tolerance

	if uA >= t0 && uA <= t1 && uB >= t0 && uB <= t1 {
		return true
	}
	return false
}

func LinesIntersect(l1, l2 Line) bool {
	return LinesIntersectTol(l1, l2, 0)
}

// GetTileVecs returns the vectors which make up the four corners of this tile
func GetTileVecs(tile *Tile) []Vec2 {
	return []Vec2{
		{float64(tile.Position.X) + tilePoints[0].X, float64(tile.Position.Y) + tilePoints[0].Y},
		{float64(tile.Position.X) + tilePoints[1].X, float64(tile.Position.Y) + tilePoints[1].Y},
		{float64(tile.Position.X) + tilePoints[2].X, float64(tile.Position.Y) + tilePoints[2].Y},
		{float64(tile.Position.X) + tilePoints[3].X, float64(tile.Position.Y) + tilePoints[3].Y},
	}
}

// GetTileLines returns the lines that make up the square which is this tile
func GetTileLines(tile *Tile) []Line {
	vecs := GetTileVecs(tile)

	return []Line{
		{vecs[tileLines[0][0]], vecs[tileLines[0][1]]},
		{vecs[tileLines[1][0]], vecs[tileLines[1][1]]},
		{vecs[tileLines[2][0]], vecs[tileLines[2][1]]},
		{vecs[tileLines[3][0]], vecs[tileLines[3][1]]},
	}
}

func LineIntersectsTile(line Line, tile *Tile) bool {
	tileLines := GetTileLines(tile)

	for _, tileLine := range tileLines {
		if LinesIntersect(line, tileLine) {
			return true
		}
	}

	return false
}

func LineIntersectsTileTol(line Line, tile *Tile, tolerance float64) bool {
	tileLines := GetTileLines(tile)

	for _, tileLine := range tileLines {
		if LinesIntersectTol(line, tileLine, tolerance) {
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
