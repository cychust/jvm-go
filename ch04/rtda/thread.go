package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), //todo edit
	}
}
func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) PushFrame(frame *Frame) {

}

func (self *Thread) PopFrame() *Frame {

}

func (self *Thread) currentFrame() *Frame {

}
