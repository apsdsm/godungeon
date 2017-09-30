package fakes

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/imposter"
)

type FakeActor struct {
	imposter.Fake

	MoveToTileReturns      MoveToTileRet
	GetRelativeTileReturns GetRelativeTileRet
}

// NewActor returns a new fake actor_controller
func NewActor() *FakeActor {
	f := FakeActor{}
	return &f
}

type MoveSig struct {
	Direction game.Direction
}

func (f *FakeActor) Move(direction game.Direction) {
	f.SetCall("Move", MoveSig{direction})
}

type MoveToTileSig struct {
	Tile *game.Tile
}

type MoveToTileRet struct {
	Error error
}

func (f *FakeActor) MoveToTile(tile *game.Tile) error {
	f.SetCall("MoveToTile", MoveToTileSig{tile})
	return f.MoveToTileReturns.Error
}
