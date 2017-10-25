package fakes

import (
	"github.com/apsdsm/godungeon/input"
	"github.com/apsdsm/imposter"
)

type FakeInputHandler struct {
	imposter.Fake
	GetEventsRet input.Events
}

func (f *FakeInputHandler) Events() input.Events {
	f.SetCall("Events")
	return f.GetEventsRet
}
