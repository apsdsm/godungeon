package game_test

import (
	. "github.com/apsdsm/godungeon/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {

	Describe("LineIntersectsTile method", func() {
		Context("given a tile with intersecting lines", func() {
			It("reports the intersections", func() {

				//   l1     l2
				//     \    |
				//       \+---+
				//   l3 --|2,2|--
				//        +---+
				//          | \
				//

				tile := Tile{
					Position: NewPosition(2, 2),
				}

				line1 := Line{
					Vec2{0, 0},
					Vec2{3, 3},
				}

				line2 := Line{
					Vec2{2, 0},
					Vec2{2, 4},
				}

				line3 := Line{
					Vec2{0, 2},
					Vec2{4, 2},
				}

				Expect(LineIntersectsTile(line1, &tile)).To(BeTrue())
				Expect(LineIntersectsTile(line2, &tile)).To(BeTrue())
				Expect(LineIntersectsTile(line3, &tile)).To(BeTrue())
			})
		})

		Context("given a tile with lines that do no intersect", func() {
			It("does not report intersection", func() {
				//                 l2
				//      l1 ---------|---
				//           +---+  |
				//  l3 ---   |2,2|  |
				//           +---+  |
				//                  |
				//

				tile := Tile{
					Position: NewPosition(2, 2),
				}

				line1 := Line{
					Vec2{0, 0},
					Vec2{4, 0},
				}

				line2 := Line{
					Vec2{3, 0},
					Vec2{3, 4},
				}

				line3 := Line{
					Vec2{0, 2},
					Vec2{1, 2},
				}

				Expect(LineIntersectsTile(line1, &tile)).To(BeFalse())
				Expect(LineIntersectsTile(line2, &tile)).To(BeFalse())
				Expect(LineIntersectsTile(line3, &tile)).To(BeFalse())
			})
		})
	})

	Describe("LinesIntersect method", func() {
		Context("given two intersecting lines", func() {
			It("reports lines as intersecting", func() {

				//  l3
				//    \  l2
				//      \|
				// l1 ---+-----
				//       | \
				//       |   \

				l1 := Line{
					Vec2{0, 1},
					Vec2{3, 1},
				}

				l2 := Line{
					Vec2{1, 0},
					Vec2{1, 3},
				}

				l3 := Line{
					Vec2{0, 0},
					Vec2{3, 3},
				}

				Expect(LinesIntersect(l1, l2)).To(BeTrue())
				Expect(LinesIntersect(l1, l3)).To(BeTrue())
				Expect(LinesIntersect(l2, l3)).To(BeTrue())
			})
		})
		Context("given two lines that do not intersect", func() {
			It("report lines as not intersecting", func() {

				// l2 ------------
				// l1 ------------
				// l3 \
				//      \
				//        \
				//

				l1 := Line{
					Vec2{0, 1},
					Vec2{3, 1},
				}

				l2 := Line{
					Vec2{0, 0},
					Vec2{3, 0},
				}

				l3 := Line{
					Vec2{0, 3},
					Vec2{3, 6},
				}

				Expect(LinesIntersect(l1, l2)).To(BeFalse())
				Expect(LinesIntersect(l2, l3)).To(BeFalse())
			})
		})
	})

	Describe("LinesIntersectT", func() {
		Context("Given two lines that are tangent at one point", func() {
			It("Returns false if the threshold is above 0", func() {

				// l1 -----
				// l2 \
				//      \
				//        \
				//

				l1 := Line{
					Vec2{0, 0},
					Vec2{3, 0},
				}

				l2 := Line{
					Vec2{0, 0},
					Vec2{3, 3},
				}

				Expect(LinesIntersectTol(l1, l2, 0.05)).To(BeFalse())
				Expect(LinesIntersectTol(l1, l2, 0)).To(BeTrue())
			})
		})
	})

	Describe("TilesInRange method", func() {
		Context("Given a map of tiles", func() {
			It("Returns tiles that are in range of given tile", func() {

				// X X X X X
				// O X X X X
				// O O X X X
				// O O O X X
				// S O O O X

				tiles := make([][]Tile, 5)

				for x := 0; x < 5; x++ {
					tiles[x] = make([]Tile, 5)
					for y := 0; y < 5; y++ {
						tiles[x][y] = Tile{Position: NewPosition(x, y)}
					}
				}

				start := tiles[0][4]
				inRange := TilesInRange(&start, tiles, 3)

				Expect(len(inRange)).To(Equal(10))
				Expect(inRange).To(ContainElement(&tiles[0][1]))
				Expect(inRange).To(ContainElement(&tiles[0][2]))
				Expect(inRange).To(ContainElement(&tiles[0][3]))
				Expect(inRange).To(ContainElement(&tiles[0][4]))
				Expect(inRange).To(ContainElement(&tiles[1][2]))
				Expect(inRange).To(ContainElement(&tiles[1][3]))
				Expect(inRange).To(ContainElement(&tiles[1][4]))
				Expect(inRange).To(ContainElement(&tiles[2][3]))
				Expect(inRange).To(ContainElement(&tiles[2][4]))
				Expect(inRange).To(ContainElement(&tiles[3][4]))
			})
		})
	})
})
