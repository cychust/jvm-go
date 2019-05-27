package classfile

import "fmt"

/*
CONSTANT_Class_info{
	u1 tag;
	u2 name_index;		//索引值 指向常量池中一个CONSTANT_Utf8_info
}
 */

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16 //指向 Constant_Utf8_info 常量
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()

	fmt.Printf("\t\tname_index:%d\n", self.nameIndex)
	fmt.Printf("\t}\n")
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
