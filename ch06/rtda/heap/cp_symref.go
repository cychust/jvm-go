package heap

type SymRef struct {
	cp        *ConstantPool //存放符号引用所在的运行时常量池指针，这样就可以通过符号引用访问到运行时常量池
	className string        //
	class     *Class
}

func (self SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolvedClassRef()
	}
	return self.class
}

func (self *SymRef) resolvedClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
