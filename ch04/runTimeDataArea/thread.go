package runTimeDataArea

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *Frame) {
	return self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread)  CurrentFrame() *Frame{
	return self.stack.top()
}