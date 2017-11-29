package game

type Damage struct {
	Dp   int
	Type DamageType
}

type DamageType = int

// This table defines the types of damage.
const (
	DamageHit = iota
	DamageMiss
)
