package bits

func cellToPowerOfTwo(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	r := x - 1
	r |= r >> 1  // 2
	r |= r >> 2  // 4
	r |= r >> 4  // 8
	r |= r >> 8  // 16
	r |= r >> 16 // 32
	r |= r >> 32 // 64
	return r + 1
}
