package fakes

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/imposter"
)

type FakeActorController struct {
	imposter.Fake
	MoveError   error
	AttackError error
}

type MoveSig struct {
	Actor *game.Actor
	Tile  *game.Tile
}

func (f *FakeActorController) Move(actor *game.Actor, tile *game.Tile) error {
	f.SetCall("Move", MoveSig{actor, tile})
	return f.MoveError
}

type AttackSig struct {
	Actor  *game.Actor
	Target *game.Actor
}

func (f *FakeActorController) Attack(actor *game.Actor, target *game.Actor) error {
	f.SetCall("Attack", AttackSig{actor, target})
	return f.AttackError
}
