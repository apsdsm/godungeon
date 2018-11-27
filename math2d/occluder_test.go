package math2d_test

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/math2d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const X = false
const O = true
const V = true

var _ = Describe("Occluder", func() {

	//     0   1   2
	// 0  [x] [x] [x]    ->   [v] [v] [v]
	// 1  [p] [ ] [ ]    ->   [v] [v] [v]
	// 2  [x] [x] [x]    ->   [v] [v] [v]
	It("finds visible tiles in a hallway", func() {

		grid := makeGrid([]bool{
			X, X, X,
			O, O, O,
			X, X, X,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][1], grid)

		checkVis(grid, []bool{
			O, O, O,
			O, O, O,
			O, O, O,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [ ] [x] [ ]    ->   [v] [v] [ ]
	// 1  [p] [x] [ ]    ->   [v] [v] [ ]
	// 2  [ ] [x] [ ]    ->   [v] [v] [ ]
	It("cannot see past walls", func() {
		grid := makeGrid([]bool{
			O, X, O,
			O, X, O,
			O, X, O,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][0], grid)

		checkVis(grid, []bool{
			V, V, X,
			V, V, X,
			V, V, X,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [ ] [x] [ ]    ->   [v] [v] [ ]
	// 1  [p] [x] [ ]    ->   [v] [v] [ ]
	// 2  [ ] [ ] [ ]    ->   [v] [v] [v]
	It("can see around corners", func() {
		grid := makeGrid([]bool{
			O, X, O,
			O, X, O,
			O, O, O,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][1], grid)

		checkVis(grid, []bool{
			V, V, X,
			V, V, X,
			V, V, X,
		}, 3, 3)
	})
})

func checkVis(in [][]game.Tile, visMap []bool, xdim, ydim int) {
	for y := 0; y < ydim; y++ {
		for x := 0; x < xdim; x++ {
			Expect(in[x][y].Visible).To(Equal(visMap[y*xdim+x]))
		}
	}
}

func makeGrid(in []bool, xdim, ydim int) [][]game.Tile {

	// make a grid using the above dims
	grid := make([][]game.Tile, xdim)

	for x := range grid {
		grid[x] = make([]game.Tile, ydim)
	}

	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			grid[x][y] = game.Tile{
				Walkable: in[y*xdim+x],
				Visible:  false,
				Position: game.NewPosition(x, y),
			}
		}
	}

	return grid
}
