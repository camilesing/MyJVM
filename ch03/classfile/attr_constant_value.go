package classfile

//ConstantValue是定长属性，只会出现在field_info结构中，用于表示常量表达式的值（详见Java语言规范的15.28节）。
//ConstantValue_attribute {
//u2 attribute_name_index;
//u4 attribute_length;
//u2 constantvalue_index;
//}
type ConstantValueAttribute struct {
	ConstantValueAttribute uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.ConstantValueAttribute = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.ConstantValueAttribute
}