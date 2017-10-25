package game_test

import (
	. "github.com/apsdsm/godungeon/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dungeon", func() {
	Describe("NewDungeon", func() {
		It("makes a new dungeon at specified size", func() {
			dungeon := NewDungeon(2, 3)
			Expect(dungeon.Width).To(Equal(2))
			Expect(dungeon.Height).To(Equal(3))
			Expect(len(dungeon.Tiles)).To(Equal(2))

			for i := 0; i < len(dungeon.Tiles); i++ {
				Expect(len(dungeon.Tiles[i])).To(Equal(3))
			}
		})
	})
	Describe("At", func() {
		Context("provided with valid coords", func() {
			It("returns pointer to tile at coordinates", func() {
				dungeon := NewDungeon(2, 2)
				Expect(dungeon.At(1, 1)).To(Equal(&dungeon.Tiles[1][1]))
			})
		})
	})
	Describe("InBounds", func() {
		Context("provided with vaid coords", func() {
			It("returns true", func() {
				dungeon := NewDungeon(2, 2)
				Expect(dungeon.InBounds(1, 1)).To(BeTrue())
			})
		})
		Context("provided with invalid coords", func() {
			It("returns false", func() {
				dungeon := NewDungeon(2, 2)
				Expect(dungeon.InBounds(-1, -1)).To(BeFalse()) // under range
				Expect(dungeon.InBounds(2, 2)).To(BeFalse())   // over range
			})
		})
	})
})
