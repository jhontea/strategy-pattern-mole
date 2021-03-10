package main

import "fmt"

// SkillForkedLightning ...
type SkillForkedLightning struct {
}

// NewSkillForkedLightning ...
func NewSkillForkedLightning() Skill1Behaviour {
	return &SkillForkedLightning{}
}

// Cast1 ...
func (s *SkillForkedLightning) Cast1() {
	fmt.Println("cast skill 1, forked lightning with many target in range")
}
