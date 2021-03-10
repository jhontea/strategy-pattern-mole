package main

import "fmt"

// SkillThunderWrath ...
type SkillThunderWrath struct {
}

// NewSkillThunderWrath ...
func NewSkillThunderWrath() SkillUltimateBehaviour {
	return &SkillThunderWrath{}
}

// CastUltimate ...
func (s *SkillThunderWrath) CastUltimate() {
	fmt.Println("cast skill ultimate, thunder wrath 3x")
}
