package classfile

//Deprecated_attribute {
//u2 attribute_name_index;
//u4 attribute_length;
//}
//Java中的@Deprecated，用于标记方法已经弃用
type DeprecatedAttribute struct{ MarkerAttribute }

//Synthetic_attribute {
//u2 attribute_name_index;
//u4 attribute_length;
//}
//用于标记源文件中不存在、由编译器生成的类成员
type SyntheticAttribute struct{  MarkerAttribute }

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//read nothing
}
