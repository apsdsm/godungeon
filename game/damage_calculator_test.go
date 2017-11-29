package game_test

import (
	. "github.com/apsdsm/godungeon/game"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DamageCalculator", func() {

	It("calculates base damage", func() {
		calc := DamageCalculator{}
		calc.SetSeed(123456)

		a := Attack{}
		a.MinDamage = 10
		a.MaxDamage = 20
		a.ChanceToHit = 90
		// total raw damage = 17

		d := Defence{}
		d.DamageCutRate = 10
		d.DamageCutCeil = 10
		// cuts 2 pt of damage

		damage := calc.CalcDamage(a, d)

		Expect(damage.Dp).To(Equal(15))
		Expect(damage.Type).To(Equal(DamageHit))
	})

	It("limits damage cut", func() {
		calc := DamageCalculator{}
		calc.SetSeed(123456)

		a := Attack{}
		a.MinDamage = 10
		a.MaxDamage = 20
		a.ChanceToHit = 90
		// total raw damage = 17

		d := Defence{}
		d.DamageCutRate = 100
		d.DamageCutCeil = 2
		// cuts 2 pt of damage (even though could potentially cut all damage)

		damage := calc.CalcDamage(a, d)

		Expect(damage.Dp).To(Equal(15))
		Expect(damage.Type).To(Equal(DamageHit))
	})

	It("returns a miss if hit does not connect", func() {
		calc := DamageCalculator{}
		calc.SetSeed(123456)

		a := Attack{}
		a.MinDamage = 10
		a.MaxDamage = 20
		a.ChanceToHit = 0
		// total raw damage = NA

		d := Defence{}
		d.DamageCutRate = 100
		d.DamageCutCeil = 2
		// cuts 2 pt of damage (even though could potentially cut all damage)

		damage := calc.CalcDamage(a, d)

		Expect(damage.Dp).To(Equal(0))
		Expect(damage.Type).To(Equal(DamageMiss))
	})
})
