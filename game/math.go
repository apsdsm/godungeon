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

func TDeg(from, to *Tile) float64 {
	b := Vec2{float64(to.Position.X), float64(to.Position.Y)}
	a := Vec2{float64(from.Position.X), float64(from.Position.Y)}
	ba := Vec2Sub(b, a)
	norm := Vec2Nor(ba)
	deg := math.Atan2(norm.x, norm.y)*180/math.Pi + 180.0

	return deg
}

func TVis(from, to *Tile) bool {
	check := from

	for check != to {
		deg := TDeg(check, to)

		// 1. find the angle that describes a line from the center of FROM to the closest corner
		//    of TO
		//
		// 2. given that triangle, when ADJ is length 0, the height of the OPP side will be 0.
		// 3. add 1.0 to the length of the ADJ and see the new hight of OPP side. this lets us
		// know which tile to check for visibility.

		//
		//           | <- opp
		//  +--------+
		//    adj          <- in this case we get tan(theta)=opp/adj
		//                    or opp x tan(theta) = adj
		//                    we know theta, so we can get the heights at multiple points
		//                    on the triangle. this allows us to chek which grid square to check
		//                    that is, if the height is 1.3, and the y cord of the starting point
		//                    is 10, then we need to chec if the tile in position 12 obstructs the
		//                    view.

		//   45   360   315
		//   90   NAN   270
		//  135   180   225

		// if the vector points to a cardinal direction, use that direction as the next check tile,
		// otherwise be a little more permissive in choosing the next check tile.

		// this monstrosity below is one of those times I wish go had a more succint
		// syntax for if/else...

		// vague corners:
		// for angles that are between cardinal points (e.g., 20degs), we have slightly more involved logic to
		// to check where the next path in the tile is. We try to be permissive in letting the line of sight
		// continue so long as it is not directly blocked in by a corner.
		//
		//   ····t··
		//   ··p#··· <- in this case t is not quite 45 degress. If we try move E the line of sight stops.
		//   #######    however since the NE tile is walkable, we use that tile instead.
		//
		//   ###·t··
		//   ··p#··· <- in this case E is blocked. NE is free, but N is blocked, so t is not visible.
		//   #######
		//
		// cardinal corners:
		// when we encounter a hard diagonal (e.g., NW), we check the surrounding tiles before deciding if we
		// can see in that direction or not. If both N and W are not walkable, then we only pass the NW tile
		// back if it is also not walkable. This allows us to fill in corners of buildings on a map without
		// also allowing a player to peek into another room through a corner.
		//
		//  t###
		//  #···
		//  #·p·   <- player cannot see t
		//
		//  ####
		//  #···
		//  #·p·   <- player can see the NW corner

		if 0 < deg && deg < 45 {
			if check.Neighbors[N].Walkable {
				check = check.Neighbors[N]
			} else if check.Neighbors[W].Walkable {
				check = check.Neighbors[NW]
			} else {
				return false
			}

		} else if deg == 45 {
			if !check.Neighbors[N].Walkable && !check.Neighbors[W].Walkable {
				if check.Neighbors[NW].Walkable {
					return false
				}
			}
			check = check.Neighbors[NW]

		} else if deg < 90 {
			if check.Neighbors[W].Walkable {
				check = check.Neighbors[W]
			} else if check.Neighbors[N].Walkable {
				check = check.Neighbors[NW]
			} else {
				return false
			}

		} else if deg == 90 {
			check = check.Neighbors[W]

		} else if deg < 135 {
			if check.Neighbors[W].Walkable {
				check = check.Neighbors[W]
			} else if check.Neighbors[S].Walkable {
				check = check.Neighbors[SW]
			} else {
				return false
			}

		} else if deg == 135 {
			if !check.Neighbors[S].Walkable && !check.Neighbors[W].Walkable {
				if check.Neighbors[SW].Walkable {
					return false
				}
			}
			check = check.Neighbors[SW]

		} else if deg < 180 {
			if check.Neighbors[S].Walkable {
				check = check.Neighbors[S]
			} else if check.Neighbors[W].Walkable {
				check = check.Neighbors[SW]
			} else {
				return false
			}

		} else if deg == 180 {
			check = check.Neighbors[S]

		} else if deg < 225 {
			if check.Neighbors[S].Walkable {
				check = check.Neighbors[S]
			} else if check.Neighbors[E].Walkable {
				check = check.Neighbors[SE]
			} else {
				return false
			}

		} else if deg == 225 {
			if !check.Neighbors[S].Walkable && !check.Neighbors[E].Walkable {
				if check.Neighbors[SE].Walkable {
					return false
				}
			}
			check = check.Neighbors[SE]

		} else if deg < 270 {
			if check.Neighbors[E].Walkable {
				check = check.Neighbors[E]
			} else if check.Neighbors[S].Walkable {
				check = check.Neighbors[SE]
			} else {
				return false
			}

		} else if deg == 270 {
			check = check.Neighbors[E]

		} else if deg < 315 {
			if check.Neighbors[E].Walkable {
				check = check.Neighbors[E]
			} else if check.Neighbors[N].Walkable {
				check = check.Neighbors[NE]
			} else {
				return false
			}

		} else if deg == 315 {
			if !check.Neighbors[N].Walkable && !check.Neighbors[E].Walkable {
				if check.Neighbors[NE].Walkable {
					return false
				}
			}

			check = check.Neighbors[NE]

		} else if deg < 360 {
			if check.Neighbors[N].Walkable {
				check = check.Neighbors[N]
			} else if check.Neighbors[E].Walkable {
				check = check.Neighbors[NE]
			} else {
				return false
			}

		} else {
			check = check.Neighbors[N]
		}

		if check == nil || check != to && !check.Walkable {
			return false
		}
	}

	//debug.Log("visible")
	return true
}
