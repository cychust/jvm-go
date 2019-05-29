package constants

import (
	"jvm/ch05/instructions/base"
	"jvm/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//nothing to do
}
