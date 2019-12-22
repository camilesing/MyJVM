package classfile

type ConstantStringInfo struct {
	cp ConstantPool
	//常量池索引
	stringIndex uint16
}

func(self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) string() string {
	return self.cp.getUtf8(self.stringIndex)
}
