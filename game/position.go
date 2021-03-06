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

package game

// Position represents a single XY coord on the map
type Position struct {
	X, Y int
}

// NewPosition creates and returns a new position object
func NewPosition(x, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

// OutOfBounds returns true if position is out of bounds
func (p *Position) OutOfBounds(width, height int) bool {
	return p.X < 0 || p.Y < 0 || p.X >= width || p.Y >= height
}
