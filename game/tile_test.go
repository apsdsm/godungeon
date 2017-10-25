package game_test

import (
	. "github.com/apsdsm/godungeon/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tile", func() {
	Describe("Neighbor", func() {
		Context("provided valid direction", func() {
			It("returns pointer to neighbor", func() {
				neighbor := Tile{}
				tile := Tile{}

				tile.Neighbors[West] = &neighbor

				Expect(tile.Neighbor(West)).To(Equal(&neighbor))
			})
		})
	})
})
