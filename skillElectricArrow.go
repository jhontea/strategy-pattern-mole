package main

import "fmt"

// SkillElectricArrow ...
type SkillElectricArrow struct {
}

// NewSkillElectricArrow ...
func NewSkillElectricArrow() Skill2Behaviour {
	return &SkillElectricArrow{}
}

// Cast2 ...
func (s *SkillElectricArrow) Cast2() {
	fmt.Println("cast skill 2, electric arrow with 2 target")
}
