package input

import (
	"github.com/gdamore/tcell"
)

// TcellInputHandler collects input using Tcell
type TcellInputHandler struct {
	screen tcell.Screen
	events Events
}

// NewHandler returns a new instance of Handler satisfied by TcellInputHandler
func NewHandler(screen tcell.Screen) TcellInputHandler {
	h := TcellInputHandler{
		screen: screen,
	}

	return h
}

// Update will check tcell for new events and add them to an events object.
func (h *TcellInputHandler) Update() {
	// array of events
	h.events = Events{}
	h.events.Keys = make([]Key, 0, 10)

	polledEvent := h.screen.PollEvent()

	// if there is key input add to events
	switch e := polledEvent.(type) {
	case *tcell.EventKey:
		k := NewKey(e.Key(), e.Rune())
		h.events.Keys = append(h.events.Keys, k)
	}
}

// Events returns the events that have been collected by Update.
func (h *TcellInputHandler) Events() Events {
	return h.events
}

// HasKeyEvent returns true if the key event is registered
// @todo write test
func (h *TcellInputHandler) HasKeyEvent(key Key) bool {
	if len(h.events.Keys) == 0 {
		return false
	}

	for i := range h.events.Keys {
		if h.events.Keys[i] == key {
			return true
		}
	}

	return false
}
