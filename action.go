package BuckTheBug

import u "github.com/JohnKek/BuckTheBug/units"

func damageUnits(damage float64, units []*u.Unit) {
	for _, unit := range units {
		unit.HealPoint += damage
	}

}
