package encounters

import (
	"github.com/svenbledt/wotlk/sim/core"
	"github.com/svenbledt/wotlk/sim/encounters/icc"
	"github.com/svenbledt/wotlk/sim/encounters/naxxramas"
	"github.com/svenbledt/wotlk/sim/encounters/toc"
	"github.com/svenbledt/wotlk/sim/encounters/ulduar"
)

func init() {
	naxxramas.Register()
	ulduar.Register()
	toc.Register()
	icc.Register()
}

func AddSingleTargetBossEncounter(presetTarget *core.PresetTarget) {
	core.AddPresetTarget(presetTarget)
	core.AddPresetEncounter(presetTarget.Config.Name, []string{
		presetTarget.Path(),
	})
}
