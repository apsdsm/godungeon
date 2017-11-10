package fakes

import (
	"github.com/apsdsm/godungeon/game"
	"github.com/apsdsm/imposter"
)

func NewDamageCalculator() FakeDamageCalculator {
	f := FakeDamageCalculator{}
	return f
}

// FakeDamageCalculator is a faked version of DamageCalculator
type FakeDamageCalculator struct {
	imposter.Fake
	CalcDamageRet CalcDamageRet
}

// CalcDamageRet is the return sig for CalcDamage
type CalcDamageRet struct {
	Damage game.Damage
}

// CalcDamageSig is the input sig for CalcDamage
type CalcDamageSig struct {
	Attack  game.Attack
	Defence game.Defence
}

// CalcDamage fakes a call to method
func (f *FakeDamageCalculator) CalcDamage(attack game.Attack, defence game.Defence) game.Damage {
	f.SetCall("CalcDamage", CalcDamageSig{attack, defence})
	return f.CalcDamageRet.Damage
}
