package input

import "github.com/gdamore/tcell"

/********************************************************************************************************************/
/* This file contains a curated subset of key values from tcell https://github.com/gdamore/tcell/blob/master/key.go */
/********************************************************************************************************************/

// KeyCode is a generic value for representing keys.go, especially special keys.go (function keys.go, cursor movement keys.go, etc.).
// For normal keys.go, like  ASCII letters, this values will be set to KeyRune, and the actual key pressed will be available
// via the rune() method of the KeyEvent interface.
//
// For more information see the tcell EventKey documentation: https://godoc.org/github.com/gdamore/tcell#EventKey
type KeyCode = tcell.Key

// provide a subset of available keys
const (
	KeyRune KeyCode = iota + 256
	KeyUp
	KeyDown
	KeyRight
	KeyLeft
	KeyUpLeft
	KeyUpRight
	KeyDownLeft
	KeyDownRight
	KeyCenter
	KeyPgUp
	KeyPgDn
	KeyHome
	KeyEnd
	KeyInsert
	KeyDelete
	KeyHelp
	KeyExit
	KeyClear
	KeyCancel
	KeyPrint
	KeyPause
	KeyBacktab
)

// a Key is a structure representing a single key press
type Key struct {
	code KeyCode
	char rune
}

// NewKey creates and returns a new key object
func NewKey(code KeyCode, char rune) Key {
	k := Key{
		code,
		char,
	}

	return k
}
