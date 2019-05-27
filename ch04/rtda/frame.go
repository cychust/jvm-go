package rtda

type Frame struct {
	lower        *Frame        //用来实现链表数据结构
	localVars    LocalVars     //保存局部变量
	operandStack *OperandStack //保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
