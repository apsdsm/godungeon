package io_test

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/godungeon/io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapLoader", func() {
	It("loads a map file", func() {
		source := "../fixtures/maps/simple.json"
		dungeon := io.LoadMap(source)

		// confirm the dungeon details were loaded
		Expect(dungeon.Desc).To(Equal("simple map with all features."))
		Expect(dungeon.Width).To(Equal(29))
		Expect(dungeon.Height).To(Equal(9))
	})

	It("initializes the player", func() {
		source := "../fixtures/maps/small.json"
		dungeon := io.LoadMap(source)

		expectedPostion := game.Position{
			X: 1,
			Y: 1,
		}

		Expect(dungeon.Player.CurrentPosition).To(Equal(expectedPostion))
	})
})
