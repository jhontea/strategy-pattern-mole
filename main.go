package main

func main() {
	skill1 := NewSkillLightning()
	skill2 := NewSkillElectricArrow()
	skillUltimate := NewSkillThunder()

	eudora := NewHero(skill1, skill2, skillUltimate)

	eudora.Skill1()
	eudora.Skill2()
	eudora.SkillUltimate()

	// new event, Mayhem
	// play in Mayhem mode
	skillUltimateMayhem := NewSkillThunderWrath()

	eudora.SetSkillUltimate(skillUltimateMayhem)
	eudora.SkillUltimate()

}
