package classfile

import "fmt"

/**
CONSTANT_NameAndType_info{
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}
 */

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()

	fmt.Printf("\t\tname_index: %d\n", self.nameIndex)
	fmt.Printf("\t\tdescriptor_index: %d\n", self.descriptorIndex)
	fmt.Printf("\t}\n")
}
