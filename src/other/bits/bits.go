package bits

import (
	"math"
)

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

// 向上取整返回最接近的 2 的整数次幂
func cellToPowerOfTwo3(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	if isPowerOfTwo(x) {
		return x
	}

	if x > math.MaxInt64 {
		return math.MaxUint64
	}

	for {
		// x -= x & (-x)
		x = x & (x - 1)
		if isPowerOfTwo(x) {
			return x << 1
		}
	}
}

// 一个二进制数字左边 0 的个数
// https://www.jianshu.com/p/2c1be41f6e59
func numberOfLeadingZeros(i uint64) uint64 {
	if i == 0 {
		return 64
	}

	var n uint64
	n = 1

	// 32
	if i>>32 == 0 {
		n += 32
		i <<= 32
	}

	// 32 + 16
	if i>>48 == 0 {
		n += 16
		i <<= 16
	}

	// 32 + 16 + 8
	if i>>56 == 0 {
		n += 8
		i <<= 8
	}

	// 32 + 16 + 8 + 4
	if i>>60 == 0 {
		n += 4
		i <<= 4
	}

	// 32 + 16 + 8 + 4 + 2
	if i>>62 == 0 {
		n += 2
		i <<= 2
	}

	n -= i >> 63
	return n
}

// 向上取整返回最接近的 2 的整数次幂
// 构造出一个跟原来数字相当的 0000 1111，然后加 1 即可
func cellToPowerOfTwo4(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	if isPowerOfTwo(x) {
		return x
	}

	if x > math.MaxInt64 {
		return math.MaxUint64
	}

	// 如果是 2 的整数次幂那段没有提前返回，这里求左侧 0 的个数的时候需要使用 x - 1
	// n := numberOfLeadingZeros(x-1)
	n := numberOfLeadingZeros(x)

	// 移位的时候为啥要 n - 1 呢？因为这里想用逻辑移位(左边补0)
	// 0x7fffffffffffffff : 0111 1111 1111 1111........
	return 0x7fffffffffffffff>>(n-1) + 1
}
