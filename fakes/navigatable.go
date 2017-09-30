package fakes

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/imposter"
)

// FakeNavigatable is a fake implementation of Navigatable
type FakeNavigatable struct {
	imposter.Fake

	GetRelativeTileReturns GetRelativeTileRet
}

// NewNavigatable returns a new fake navigatable interface
func NewNavigatable() *FakeNavigatable {
	f := FakeNavigatable{}
	return &f
}

// Incoming sig for GetRelativeTile method
type GetRelativeTileSig struct {
	Tile      *game.Tile
	Direction game.Direction
}

// Outgoing sig for GetRelativeTile method
type GetRelativeTileRet struct {
	Tile  *game.Tile
	Error error
}

// GetRelativeTile fakes the interface method
func (f *FakeNavigatable) GetRelativeTile(tile *game.Tile, direction game.Direction) (*game.Tile, error) {
	f.SetCall("GetRelativeTile", GetRelativeTileSig{tile, direction})
	return f.GetRelativeTileReturns.Tile, f.GetRelativeTileReturns.Error
}
