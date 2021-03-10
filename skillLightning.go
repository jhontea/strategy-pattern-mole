package main

import "fmt"

// SkillLightning ...
type SkillLightning struct {
}

// NewSkillLightning ...
func NewSkillLightning() Skill1Behaviour {
	return &SkillLightning{}
}

// Cast1 ...
func (s *SkillLightning) Cast1() {
	fmt.Println("cast skill 1, lightning with 1 target")
}
