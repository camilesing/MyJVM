package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

//u1
func (self *ClassReader) readUint8() uint8 {
	value := self.data[0]
	self.data = self.data[1:]
	return value
}

//u2
func (self *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return value
}

//u3
func (self *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return value
}

func (self *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return value
}

func (self *ClassReader) readUint16s() []uint16 {
	size := self.readUint16()
	slice := make([]uint16, size)
	for i := range slice {
		slice[i] = self.readUint16()
	}
	return slice
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
