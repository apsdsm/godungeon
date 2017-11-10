package game

import (
	"math"
	"math/rand"
)

type DamageCalculator struct {
	seed   int64
	random *rand.Rand
}

func NewDamageCalculator() DamageCalculator {
	d := DamageCalculator{}
	d.SetSeed(0)
	return d
}

func (d *DamageCalculator) SetSeed(s int64) {
	d.seed = s
	d.random = rand.New(rand.NewSource(s))

}

func (d *DamageCalculator) CalcDamage(attack Attack, defence Defence) Damage {

	// check if hits
	isHit := int(math.Floor(100*d.random.Float64())) < attack.ChanceToHit

	if !isHit {
		return Damage{
			Dp:   0,
			Type: DamageMiss,
		}
	}

	// pre casting to float makes the code below easier to read
	hitMin := float64(attack.MinDamage)
	hitMax := float64(attack.MaxDamage)
	cutRate := float64(defence.DamageCutRate)
	cutCeil := float64(defence.DamageCutCeil)

	// total raw damage
	hit := hitMin + math.Ceil((hitMax-hitMin)*d.random.Float64())

	// amount of damage deflected
	cut := limit(math.Ceil(hit*0.01*cutRate), cutCeil)

	damage := int(hit - cut)

	return Damage{
		Dp:   damage,
		Type: DamageHit,
	}
}

func limit(a float64, limit float64) float64 {
	if a > limit {
		return limit
	}
	return a
}
