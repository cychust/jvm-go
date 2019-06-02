package heap

type Object struct {
	//todo
	class *Class
	fields Slots
}

func newObject(class *Class) *Object  {
	return &Object{
		class:class,
		fields:newSlots(class.instantSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

