package constants

import (
	"jvm/ch06/instructions/base"
	"jvm/ch06/rtda"
)

type BIPUSH struct {
	//push a single byte constant value(-128 ~ 127) to the top of stack
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}
