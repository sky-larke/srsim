package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   20.0,
	})

	state.EndAttack()

	// Talent
	c.engine.ModifyCurrentGaugeCost(-talent[c.info.TalentLevelIndex()])
}
