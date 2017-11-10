//    Copyright 2017 Nick del Pozo
//
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

package actor_renderer

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/game"
	"github.com/gdamore/tcell"
)

// An ActorRenderer renders the entity portion of a actor_controller to a layer
type ActorRenderer struct {
	layer  *canvas.Layer
	actors *[]game.Actor
}

// NewEntityRenderer creates and returns a pointer to a a new ActorRenderer object
func New(entities *[]game.Actor, layer *canvas.Layer) *ActorRenderer {
	r := ActorRenderer{}
	r.layer = layer
	r.actors = entities
	return &r
}

// Render will send information about each entity to the assigned layer
func (r *ActorRenderer) Render() {
	style := tcell.StyleDefault
	for _, t := range *r.actors {

		// choose the color for the actor based on its state
		if t.IsPlayer {
			style = style.Foreground(game.White)
		} else {
			style = style.Foreground(game.Green)
		}

		at := t.Tile.Position
		r.layer.At(at.X, at.Y).Set(
			t.Rune,
			style,
		)
	}
}
