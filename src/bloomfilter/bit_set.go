package bloomfilter

// Bits ...
// arrar[0] [0-63]
// arrar[1] [64-127]
type Bits struct {
	length   uint64 // bit array 所表示的有效二进制的长度
	array    []uint64
	arrayLen uint64
}

func newBitSet(length uint64) *Bits {
	arrayLen := (length + 63) / 64
	array := make([]uint64, arrayLen)

	s := &Bits{
		length:   length,
		array:    array,
		arrayLen: arrayLen,
	}
	return s
}

func (s *Bits) get(index uint64) bool {
	if index >= s.length {
		return false
	}

	arrayIdx := index / 64
	bitIdx := index % 64
	bit := uint64(1 << bitIdx)

	value := s.array[arrayIdx]
	return (value & bit) != 0
}

func (s *Bits) set(index uint64) bool {
	if index >= s.length {
		return false
	}

	arrayIdx := index / 64
	bitIdx := index % 64
	bit := uint64(1 << bitIdx)

	oldValue := s.array[arrayIdx]
	s.array[arrayIdx] = oldValue | bit
	return (oldValue & bit) == 0
}
