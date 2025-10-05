//    Copyright 2016 Nick del Pozo
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

package canvas

import (
	"github.com/gdamore/tcell"
)

// Screen bridges methods used in a tcell screen to methods used by the tools. This allows for easier faking.
type Screen interface {
	SetContent(x int, y int, mainc rune, combc []rune, style tcell.Style)
	Size() (int, int)
	Show()
	Fill(r rune, s tcell.Style)
}
