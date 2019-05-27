package rtda

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			//size:  maxStack,
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

