package base

import "jvm/ch06/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) //fetch operand from byte code
	Execute(frame *rtda.Frame)            // execute instructions
}

type NoOperandsInstruction struct {
	//empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	//local var
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
