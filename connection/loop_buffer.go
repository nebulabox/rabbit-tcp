package connection

type LoopByteBuffer struct {
	buffer []byte
	head   int
	tail   int
}

func NewLoopBuffer(size uint32) LoopByteBuffer {
	buffer := make([]byte, size)
	return LoopByteBuffer{
		buffer: buffer,
	}
}

func (lb *LoopByteBuffer) OverWrite(data []byte) {
	if len(lb.buffer) < len(data) {
		lb.buffer = make([]byte, len(data))
	}
	n := copy(lb.buffer, data)
	lb.head = 0
	lb.tail = n
}

func (lb *LoopByteBuffer) Read(data []byte) int {
	n := len(data)
	if n > lb.tail-lb.head {
		n = lb.tail - lb.head
	}
	copy(data, lb.buffer[lb.head:lb.tail])
	lb.head += n
	return n
}

func (lb *LoopByteBuffer) Empty() bool {
	return lb.tail-lb.head == 0
}