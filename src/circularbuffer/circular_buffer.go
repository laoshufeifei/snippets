package circularbuffer

import "time"

// 简单的实现一个环形缓冲区，最小的 capability 为 4
// 如果是一个读一个写，不用加锁；如果需要扩容功能则需要加锁(用户保证)
// 优化策略：将 sleep 换成 条件变量

// Ring ...
type Ring struct {
	capability uint32
	buffer     []byte
	readPos    uint32
	writePos   uint32
	isEOF      bool
}

// New  ...
func New(cap int) *Ring {
	capability := uint32((cap + 3) & ^3)
	return &Ring{
		capability: capability,
		buffer:     make([]byte, capability),
	}
}

// Capability ...
func (r *Ring) Capability() uint32 {
	return r.capability
}

// IsEmpty ...
func (r *Ring) IsEmpty() bool {
	// return r.freeSize() == r.capability
	return r.writePos == r.readPos
}

// IsFull ...
func (r *Ring) IsFull() bool {
	// return r.writePos-r.readPos == r.capability
	return r.freeSize() == 0
}

// IsEOF ...
func (r *Ring) IsEOF() bool {
	return r.isEOF
}

// SetEOF ...
func (r *Ring) SetEOF() {
	r.isEOF = true
}

// Read 读到数据了就可以返回
// 如果已经 EOF, 必须读完
// 也就是说 return 长度大于 0 表示还可以继续读，否则表示结束
func (r *Ring) Read(size uint32) []byte {
	var results []byte
	for !r.IsEmpty() {
		results = r.readOnce(size)
		if len(results) > 0 {
			break
		}

		// readOnce return 0
		if r.IsEOF() && r.IsEmpty() {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	return results
}

// Write 必须一次写完，否则阻塞
func (r *Ring) Write(data []byte) uint32 {
	writedSize, allSize := uint32(0), uint32(len(data))
	for writedSize < allSize {
		if r.IsEOF() {
			break
		}

		offset := r.writeOnce(data[writedSize:])
		writedSize += offset
		if offset == 0 {
			time.Sleep(10 * time.Millisecond)
		}
	}
	return writedSize
}

// Expand 扩容一倍，注意用户要加锁
func (r *Ring) Expand() {
	oldCapacity, oldBufferSize := r.Capability(), r.bufferSize()
	newCapability := oldCapacity << 1
	newBuffer := make([]byte, newCapability)

	if oldBufferSize > 0 {
		readOffset, writeOffset := r.readOffset(), r.writeOffset()
		if writeOffset <= readOffset {
			length := uint32(copy(newBuffer, r.buffer[readOffset:oldCapacity]))
			copy(newBuffer[length:], r.buffer[:writeOffset])
		} else {
			copy(newBuffer, r.buffer[readOffset:writeOffset])
		}
	}

	r.readPos = 0
	r.writePos = oldBufferSize
	r.capability = newCapability
	r.buffer = newBuffer
}

func (r *Ring) readOffset() uint32 {
	// r.readPos % r.capability
	return r.readPos & (r.capability - 1)
}

func (r *Ring) writeOffset() uint32 {
	// r.writePos % r.capability
	return r.writePos & (r.capability - 1)
}

// 可以读的 size
func (r *Ring) bufferSize() (freeSize uint32) {
	return r.writePos - r.readPos
}

// 可以写的 size
func (r *Ring) freeSize() (freeSize uint32) {
	freeSize = r.capability - (r.writePos - r.readPos)
	return
}

// 读一次(尽力读，不一定读满，返回值是新 make 出来的)
func (r *Ring) readOnce(size uint32) []byte {
	readSize := r.bufferSize()
	if readSize > size {
		readSize = size
	}

	readOffset := r.readOffset()
	results := make([]byte, readSize)

	// 第一段是从读指针开始向缓冲区末尾方向
	copiedSize := uint32(copy(results[:readSize], r.buffer[readOffset:r.capability]))
	r.readPos += copiedSize

	// 第二段是从缓冲区起始处读入余下的可读入数据(可能为0)
	copiedSize = uint32(copy(results[copiedSize:readSize], r.buffer))
	r.readPos += copiedSize

	return results
}

// 写一次(尽力写，不一定写完)
func (r *Ring) writeOnce(data []byte) uint32 {
	writeOffset := r.writeOffset()

	// length = min{待写入数据长度, 缓冲区长度 - (写指针 - 读指针)}
	freeSize := r.freeSize()
	length := uint32(len(data))
	if length > freeSize {
		length = freeSize
	}

	// 第一段是从写指针开始向缓冲区末尾方向
	copiedSize := uint32(copy(r.buffer[writeOffset:r.capability], data[:length]))
	r.writePos += copiedSize

	// 第二段是从缓冲区起始处写入余下的可写入数据(可能为0)
	copiedSize = uint32(copy(r.buffer, data[copiedSize:length]))
	r.writePos += copiedSize

	return length
}