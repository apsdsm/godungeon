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

package render

import (
	"github.com/apsdsm/canvas"
	"github.com/apsdsm/godungeon/game"
)

// An EntityRenderer renders the entity portion of a actor_controller to a layer
type EntityRenderer struct {
}

// NewEntityRenderer creates and returns a pointer to a a new EntityRenderer object
func NewEntityRenderer() *EntityRenderer {
	r := EntityRenderer{}
	return &r
}

// DrawPlayer draws the actor_controller to a layer
func (r *EntityRenderer) DrawPlayer(player *game.Player, layer *canvas.Layer) {
	layer.Grid[player.CurrentPosition.X][player.CurrentPosition.Y].Rune = player.Rune
}
