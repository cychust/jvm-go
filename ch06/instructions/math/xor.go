package math

import (
	"jvm/ch06/instructions/base"
	"jvm/ch06/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 ^ v1
	stack.PushLong(result)
}
