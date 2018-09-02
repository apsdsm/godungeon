package json_format

type Dungeon struct {
	Width         int
	Height        int
	Link          string
	Name          string
	Desc          string
	StartPosition Position
	Tiles         [][]Tile
	Doors         []Door
	Keys          []Key
	Mobs          []Mob
}

type Tile struct {
	Rune      rune
	Walkable  bool
	Spawn     string
	Neighbors [8]Position
}

type Position struct {
	X, Y int
}

type Door struct {
	Link   string
	Locked bool
	Key    string
	OnTry  string
}

type Key struct {
	Name string
	Link string
	Desc string
}

type Mob struct {
	Name  string
	Link  string
	Prot  string
	Rune  string
	Hp    string
	Mp    string
	Sight string
}
