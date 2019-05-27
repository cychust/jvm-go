package classfile

import "fmt"

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()

	fmt.Printf("\t\tstring_index: %d\n", self.stringIndex)
	fmt.Printf("\t}\n")
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
