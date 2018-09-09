package updaters

import (
	"github.com/apsdsm/godungeon/game"
)

type Lighting struct {
	player     *game.Actor
	actors     *[]game.Actor
	dungeon    *game.Dungeon
	lightTable map[int]float32
}

type LightingConfig struct {
	Player  *game.Actor
	Actors  *[]game.Actor
	Dungeon *game.Dungeon
}

func NewLighting(config LightingConfig) *Lighting {
	l := Lighting{}
	l.player = config.Player
	l.actors = config.Actors
	l.dungeon = config.Dungeon

	for i := 0; i < l.player.Sight; i++ {
		l.lightTable[i] = float32(i * i)
	}

	return &l
}

func (u *Lighting) Update() {
	for x := 0; x < len(u.dungeon.Tiles); x++ {
		for y := 0; y < len(u.dungeon.Tiles[x]); y++ {

			// reset brightness
			u.dungeon.Tiles[x][y].Brightness = 0

			// distance to tile
			dist := game.TDist(u.player.Tile, &u.dungeon.Tiles[x][y])

			if dist > u.player.Sight {
				u.dungeon.Tiles[x][y].Visible = false
				continue
			} else if game.TVis(u.player.Tile, &u.dungeon.Tiles[x][y]) {
				u.dungeon.Tiles[x][y].Seen = true
				u.dungeon.Tiles[x][y].Visible = true
				u.dungeon.Tiles[x][y].Brightness = 1 / u.lightTable[dist]
			}
		}
	}
}
