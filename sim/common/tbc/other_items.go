package tbc

import (
	"github.com/svenbledt/wotlk/sim/core"
	"github.com/svenbledt/wotlk/sim/core/stats"
)

func init() {
	core.NewItemEffect(30892, func(agent core.Agent) {
		for _, pet := range agent.GetCharacter().Pets {
			if pet.IsGuardian() {
				continue // not sure if this applies to guardians.
			}
			pet.PseudoStats.DamageDealtMultiplier *= 1.03
			pet.AddStat(stats.MeleeCrit, core.CritRatingPerCritChance*2)
		}
	})
}
