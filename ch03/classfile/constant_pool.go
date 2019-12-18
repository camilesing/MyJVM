package classfile

//常量池实际上也是一个表，但是有三点需要特别注意。
//第一，表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常量池的实际大小是n–1。
//第二，有效的常量池索引是1~n–1。0是无效索引，表示不指向任何常量。
//第三，CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。也就是说，如果常量池中存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些数也会变成无效索引。
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { //see note, start index with 1
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(tpye) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //see note, its need 2 index
		}
	}
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(utInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	calssInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
