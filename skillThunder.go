package main

import "fmt"

// SkillThunder ...
type SkillThunder struct {
}

// NewSkillThunder ...
func NewSkillThunder() SkillUltimateBehaviour {
	return &SkillThunder{}
}

// CastUltimate ...
func (s *SkillThunder) CastUltimate() {
	fmt.Println("cast skill ultimate, thunder 1x")
}
