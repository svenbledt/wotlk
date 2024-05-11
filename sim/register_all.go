package sim

import (
	_ "github.com/svenbledt/wotlk/sim/common"
	dpsDeathKnight "github.com/svenbledt/wotlk/sim/deathknight/dps"
	tankDeathKnight "github.com/svenbledt/wotlk/sim/deathknight/tank"
	"github.com/svenbledt/wotlk/sim/druid/balance"
	"github.com/svenbledt/wotlk/sim/druid/feral"
	restoDruid "github.com/svenbledt/wotlk/sim/druid/restoration"
	feralTank "github.com/svenbledt/wotlk/sim/druid/tank"
	_ "github.com/svenbledt/wotlk/sim/encounters"
	"github.com/svenbledt/wotlk/sim/hunter"
	"github.com/svenbledt/wotlk/sim/mage"
	holyPaladin "github.com/svenbledt/wotlk/sim/paladin/holy"
	protectionPaladin "github.com/svenbledt/wotlk/sim/paladin/protection"
	"github.com/svenbledt/wotlk/sim/paladin/retribution"
	healingPriest "github.com/svenbledt/wotlk/sim/priest/healing"
	"github.com/svenbledt/wotlk/sim/priest/shadow"
	"github.com/svenbledt/wotlk/sim/priest/smite"
	"github.com/svenbledt/wotlk/sim/rogue"
	"github.com/svenbledt/wotlk/sim/shaman/elemental"
	"github.com/svenbledt/wotlk/sim/shaman/enhancement"
	restoShaman "github.com/svenbledt/wotlk/sim/shaman/restoration"
	"github.com/svenbledt/wotlk/sim/warlock"
	dpsWarrior "github.com/svenbledt/wotlk/sim/warrior/dps"
	protectionWarrior "github.com/svenbledt/wotlk/sim/warrior/protection"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	feral.RegisterFeralDruid()
	feralTank.RegisterFeralTankDruid()
	restoDruid.RegisterRestorationDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	restoShaman.RegisterRestorationShaman()
	hunter.RegisterHunter()
	mage.RegisterMage()
	healingPriest.RegisterHealingPriest()
	shadow.RegisterShadowPriest()
	smite.RegisterSmitePriest()
	rogue.RegisterRogue()
	dpsWarrior.RegisterDpsWarrior()
	protectionWarrior.RegisterProtectionWarrior()
	holyPaladin.RegisterHolyPaladin()
	protectionPaladin.RegisterProtectionPaladin()
	retribution.RegisterRetributionPaladin()
	warlock.RegisterWarlock()
	dpsDeathKnight.RegisterDpsDeathknight()
	tankDeathKnight.RegisterTankDeathknight()
}
