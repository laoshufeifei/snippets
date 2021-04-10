package bits

func isPowerOfTwo(n uint64) bool {
	if n == 0 {
		return false
	}

	return (n & (n - 1)) == 0
}

// 向上取整返回最接近的 2 的整数次幂
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

// 向上取整返回最接近的 2 的整数次幂
func cellToPowerOfTwo2(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	if isPowerOfTwo(x) {
		return x
	}

	x |= x >> 1  // 2
	x |= x >> 2  // 4
	x |= x >> 4  // 8
	x |= x >> 8  // 16
	x |= x >> 16 // 32
	x |= x >> 32 // 64
	return x + 1
}
