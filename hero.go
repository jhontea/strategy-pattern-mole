package main

// Hero ...
type Hero struct {
	skill1        Skill1Behaviour
	skill2        Skill2Behaviour
	skillUltimate SkillUltimateBehaviour
}

// NewHero ...
func NewHero(
	skill1 Skill1Behaviour,
	skill2 Skill2Behaviour,
	skillUltimate SkillUltimateBehaviour,
) *Hero {
	return &Hero{
		skill1:        skill1,
		skill2:        skill2,
		skillUltimate: skillUltimate,
	}
}

// SetSkill1 ...
func (h *Hero) SetSkill1(skill Skill1Behaviour) {
	h.skill1 = skill
}

// SetSkill2 ...
func (h *Hero) SetSkill2(skill Skill2Behaviour) {
	h.skill2 = skill
}

// SetSkillUltimate ...
func (h *Hero) SetSkillUltimate(skill SkillUltimateBehaviour) {
	h.skillUltimate = skill
}

// Skill1 ...
func (h *Hero) Skill1() {
	h.skill1.Cast1()
}

// Skill2 ...
func (h *Hero) Skill2() {
	h.skill2.Cast2()
}

// SkillUltimate ...
func (h *Hero) SkillUltimate() {
	h.skillUltimate.CastUltimate()
}
