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

package renderers

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

// NewActorRenderer creates and returns a pointer to a a new ActorRenderer object
func NewActorRenderer(entities *[]game.Actor, layer *canvas.Layer) *ActorRenderer {
	r := ActorRenderer{}
	r.layer = layer
	r.actors = entities
	return &r
}

// Render will send information about each entity to the assigned layer
func (r *ActorRenderer) Render() {
	for _, a := range *r.actors {
		at := a.Tile.Position

		r.layer.At(at.X, at.Y).Set(
			renderRune(a),
			renderStyle(a),
		)
	}
}

// The correct rune to render for this actor. If the actor is alive it returns
// the actor's specified rune. Otherwise it returns a cross.
func renderRune(a game.Actor) rune {
	if a.IsDead {
		return '✝'
	}

	return a.Rune
}

// The style in which this actor should be rendered. Player is white. Mobs are
// green until they are dead, when they are red.
func renderStyle(a game.Actor) tcell.Style {
	style := tcell.StyleDefault

	if a.IsPlayer {
		return style.Foreground(game.White)
	}

	if hpPercentRemaining(a) <= 0.10 {
		return style.Foreground(game.Red)
	}

	return style.Foreground(game.Green)
}

// hp remaining for this actor.
func hpPercentRemaining(a game.Actor) float64 {
	return float64(a.Hp) / float64(a.MaxHp)
}
