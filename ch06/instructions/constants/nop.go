package constants

import (
	"jvm/ch06/instructions/base"
	"jvm/ch06/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//nothing to do
}
