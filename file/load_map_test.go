//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package file_test

import (
	"github.com/apsdsm/godungeon/file"

	"github.com/apsdsm/godungeon/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("MapLoader", func() {

	Context("loading a map", func() {
		var (
			source  string
			dungeon *game.Dungeon
		)

		It("loads tiles", func() {
			source = "../fixtures/maps/simple.json"
			dungeon = file.LoadMap(source)

			Expect(dungeon.Width).To(Equal(29))
			Expect(dungeon.Height).To(Equal(9))
		})

		DescribeTable("it sets correct neighbors",
			func(check int, neighbors [8]int) {

				source = "../fixtures/maps/neighbors.json"
				dungeon = file.LoadMap(source)

				// tile coords
				tile := make(map[int]*game.Tile)
				tile[-1] = nil
				tile[0] = dungeon.At(0, 0)
				tile[1] = dungeon.At(1, 0)
				tile[2] = dungeon.At(2, 0)
				tile[3] = dungeon.At(0, 1)
				tile[4] = dungeon.At(1, 1)
				tile[5] = dungeon.At(2, 1)
				tile[6] = dungeon.At(0, 2)
				tile[7] = dungeon.At(1, 2)
				tile[8] = dungeon.At(2, 2)

				for i, n := range neighbors {
					Expect(tile[check].Neighbors[i]).To(Equal(tile[n]))
				}
			},
			Entry("tile 0", 0, [8]int{-1, -1, 1, 4, 3, -1, -1, -1}),
			Entry("tile 1", 1, [8]int{-1, -1, 2, 5, 4, 3, 0, -1}),
			Entry("tile 2", 2, [8]int{-1, -1, -1, -1, 5, 4, 1, -1}),
			Entry("tile 3", 3, [8]int{0, 1, 4, 7, 6, -1, -1, -1}),
			Entry("tile 4", 4, [8]int{1, 2, 5, 8, 7, 6, 3, 0}),
			Entry("tile 5", 5, [8]int{2, -1, -1, -1, 8, 7, 4, 1}),
			Entry("tile 6", 6, [8]int{3, 4, 7, -1, -1, -1, -1, -1}),
			Entry("tile 7", 7, [8]int{4, 5, 8, -1, -1, -1, 6, 3}),
			Entry("tile 8", 8, [8]int{5, -1, -1, -1, -1, -1, 7, 4}),
		)

		It("initializes the player", func() {
			source = "../fixtures/maps/mobs.json"
			dungeon = file.LoadMap(source)

			Expect(dungeon.Actors[0].Name).To(Equal("player"))
			Expect(dungeon.Actors[0].Rune).To(Equal('x'))
			Expect(dungeon.Actors[0].Tile).To(Equal(dungeon.At(2, 1)))
			Expect(dungeon.Actors[0].IsPlayer).To(BeTrue())
		})

		It("initializes mobs", func() {
			source := "../fixtures/maps/mobs.json"
			dungeon = file.LoadMap(source)

			Expect(len(dungeon.Actors)).To(Equal(2))
			Expect(dungeon.Actors[1].Rune).To(Equal('m'))
			Expect(dungeon.Actors[1].Tile).To(Equal(dungeon.At(3, 2)))
			Expect(dungeon.Actors[1].IsPlayer).To(BeFalse())
			Expect(dungeon.Actors[1].Hp).To(Equal(10))
			Expect(dungeon.Actors[1].MaxHp).To(Equal(10))
			Expect(dungeon.Actors[1].Mp).To(Equal(20))
			Expect(dungeon.Actors[1].MaxMp).To(Equal(20))
		})

		It("initializes items", func() {
			source := "../fixtures/maps/mobs.json"
			dungeon = file.LoadMap(source)
		})
	})

})
