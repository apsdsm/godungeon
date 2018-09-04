package game

import (
	"fmt"
	"math"
	"strconv"

	"github.com/apsdsm/godungeon/debug"
)

// TDist returns the number of steps between two tiles (assuming a diagonal step is 2)
func TDist(from, to *Tile) int {
	xdis := math.Abs(float64(from.Position.X) - float64(to.Position.X))
	ydis := math.Abs(float64(from.Position.Y) - float64(to.Position.Y))

	return int(xdis + ydis)
}

type Vec2 struct {
	x float64
	y float64
}

func Vec2Sub(a, b Vec2) Vec2 {
	return Vec2{a.x - b.x, a.y - b.y}
}

func Vec2Add(a, b Vec2) Vec2 {
	return Vec2{a.x + b.x, a.y + b.y}
}

func Vec2MulF(a Vec2, b float64) Vec2 {
	return Vec2{a.x * b, a.y * b}
}

func Vec2Mag(a Vec2) float64 {
	return a.x*a.x + a.y*a.y
}

func Vec2Len(a Vec2) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

func Vec2Nor(a Vec2) Vec2 {
	len := Vec2Len(a)
	return Vec2{a.x / len, a.y / len}
}

func Vec2Dot(a, b Vec2) float64 {
	return a.x*b.x + a.y*b.y
}

type pftile struct {
	tile   *Tile
	dist   int
	parent *Tile
	open   bool
}

func TVis(from, to *Tile) bool {

	b := Vec2{float64(to.Position.X), float64(to.Position.Y)}
	check := from

	count := 0

	for check != to {

		debug.Log(strconv.Itoa(count))
		count++

		a := Vec2{float64(check.Position.X), float64(check.Position.Y)}
		ba := Vec2Sub(b, a) // target b to a
		norm := Vec2Nor(ba)
		deg := math.Atan2(norm.x, norm.y)*180/math.Pi + 180.0

		debug.Log(fmt.Sprintf("%f", deg))

		//   45   360   315
		//   90   NAN   270
		//  135   180   225

		// put checks in for non-specific angles like between 45 and 0
		// so that it goes E unless that's blocked in which case it goes NE

		if deg < 45 || deg > 315 {
			check = check.Neighbors[N]
			//debug.Log("N")
		} else if deg == 315 {
			check = check.Neighbors[NE]
			//debug.Log("NE")
		} else if deg < 315 && deg > 225 {
			check = check.Neighbors[E]
			//debug.Log("E")
		} else if deg == 225 {
			check = check.Neighbors[SE]
			//debug.Log("SE")
		} else if deg < 225 && deg > 135 {
			check = check.Neighbors[S]
			//debug.Log("S")
		} else if deg == 135 {
			check = check.Neighbors[SW]
			//debug.Log("SW")
		} else if deg < 135 && deg > 45 {
			check = check.Neighbors[W]
			//debug.Log("W")
		} else {
			check = check.Neighbors[NW]
			//debug.Log("NW")
		}

		if check == nil {
			//debug.Log("nil - not visible")
			return false

		}

		if check != to && !check.Walkable {
			//debug.Log("obstructed - not visible")
			return false
		}
	}

	//debug.Log("visible")
	return true
}
