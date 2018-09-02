package updaters

import (
	"github.com/apsdsm/godungeon/game"
)

// MobAi controls mobs in the game
type MobAi struct {
	player *game.Actor
	mobs   *[]game.Actor
}

// MobAiConfig stores the config info for the constructor
type MobAiConfig struct {
	Player *game.Actor
	Mobs   *[]game.Actor
}

// NewMobAi creates and returns a new mob ai
func NewMobAi(config MobAiConfig) MobAi {
	m := MobAi{}
	m.player = config.Player
	m.mobs = config.Mobs

	return m
}

// Update mobs in the level
func (u *MobAi) Update() {
	//debug.Log("updating mob ai")

	for i := 0; i < len(*u.mobs); i++ {

		if (*u.mobs)[i].IsPlayer {
			continue
		}

		dist := game.TDist((*u.mobs)[i].Tile, u.player.Tile)

		if dist < (*u.mobs)[i].Sight {
			//debug.Log("I see you!")
		} else {
			//debug.Log("I can't see you")
		}
	}
}
