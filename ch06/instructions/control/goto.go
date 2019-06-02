package control

import (
	"jvm/ch06/instructions/base"
	"jvm/ch06/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

