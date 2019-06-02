package rtda

import "jvm/ch06/rtda/heap"

type Frame struct {
	lower        *Frame        //用来实现链表数据结构
	localVars    LocalVars     //保存局部变量
	operandStack *OperandStack //保存操作数栈指针
	thread       *Thread
	nextPC       int
	method       *heap.Method
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    NewLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Method() *heap.Method {
	return self.method
}
