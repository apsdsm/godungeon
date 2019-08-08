package math2d_test

import (
	"fmt"
	"testing"

	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/math2d"
	"github.com/stretchr/testify/assert"
)

const X = false
const O = true

func TestFindVisibleTiles2(t *testing.T) {
	t.Run("it finds visible tiles in a hallway", func(t *testing.T) {
		grid := makeGrid([]bool{
			X, X, X,
			O, O, O,
			X, X, X,
		}, 3, 3)

		math2d.FindVisibleTiles2(&grid[0][1], grid)

		checkVis(t, grid, []bool{
			O, O, O,
			O, O, O,
			O, O, O,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [x] [x] [x]    ->   [v] [v] [v]
	// 1  [p] [ ] [ ]    ->   [v] [v] [v]
	// 2  [x] [x] [x]    ->   [v] [v] [v]
	t.Run("finds visible tiles in a hallway", func(t *testing.T) {

		grid := makeGrid([]bool{
			X, X, X,
			O, O, O,
			X, X, X,
		}, 3, 3)

		math2d.FindVisibleTiles2(&grid[0][1], grid)

		checkVis(t, grid, []bool{
			O, O, O,
			O, O, O,
			O, O, O,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [ ] [x] [ ]    ->   [v] [v] [ ]
	// 1  [p] [x] [ ]    ->   [v] [v] [ ]
	// 2  [ ] [x] [ ]    ->   [v] [v] [ ]
	t.Run("cannot see past walls", func(t *testing.T) {
		grid := makeGrid([]bool{
			O, X, O,
			O, X, O,
			O, X, O,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][0], grid)

		checkVis(t, grid, []bool{
			O, O, X,
			O, O, X,
			O, O, X,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [ ] [x] [ ]    ->   [v] [v] [ ]
	// 1  [p] [x] [ ]    ->   [v] [v] [ ]
	// 2  [ ] [ ] [ ]    ->   [v] [v] [v]
	t.Run("can see around corners", func(t *testing.T) {
		grid := makeGrid([]bool{
			O, X, O,
			O, X, O,
			O, O, O,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][1], grid)

		checkVis(t, grid, []bool{
			O, O, X,
			O, O, X,
			O, O, O,
		}, 3, 3)
	})

	//     0   1   2
	// 0  [x] [x] [x]    ->   [v] [v] [v]
	// 1  [ ] [ ] [x]    ->   [v] [v] [v]
	// 2  [p] [ ] [x]    ->   [v] [v] [v]
	t.Run("can see into corners", func(t *testing.T) {
		grid := makeGrid([]bool{
			X, X, X,
			O, O, X,
			O, O, X,
		}, 3, 3)

		math2d.FindVisibleTiles(&grid[0][2], grid)

		checkVis(t, grid, []bool{
			O, O, O,
			O, O, O,
			O, O, O,
		}, 3, 3)
	})

}

// compare a tile map against a visiblity map
func checkVis(t *testing.T, in [][]game.Tile, visMap []bool, xdim, ydim int) {

	errs := make([]string, 0)

	fmt.Print("diff >  expc >  actl \n")

	for y := 0; y < ydim; y++ {
		for x := 0; x < xdim; x++ {
			assert.Equal(t, in[x][y].Visible, visMap[y*xdim+x])

			if in[x][y].Visible != visMap[y*xdim+x] {
				errs = append(errs, fmt.Sprintf("err at - %d, %d\n", x, y))
				fmt.Print("-")
			} else {
				fmt.Print("+")
			}
		}

		fmt.Print("  >  ")

		for x := 0; x < xdim; x++ {
			if visMap[y*xdim+x] {
				fmt.Print("O")
			} else {
				fmt.Print("X")
			}
		}

		fmt.Print("  >  ")

		for x := 0; x < xdim; x++ {
			if in[x][y].Visible {
				fmt.Print("O")
			} else {
				fmt.Print("X")
			}
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")
	for _, err := range errs {
		fmt.Println(err)
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
