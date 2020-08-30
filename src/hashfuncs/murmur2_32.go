package hashfuncs

// https://sites.google.com/site/murmurhash/
// https://github.com/aviddiviner/go-murmur/blob/master/murmur2.go

// Mixing constants; generated offline.
const (
	M = 0x5bd1e995
	R = 24
)

// 32-bit mixing function.
func mmix(h uint32, k uint32) (uint32, uint32) {
	k *= M
	k ^= k >> R
	k *= M

	h *= M
	h ^= k
	return h, k
}

// The original MurmurHash2 32-bit algorithm by Austin Appleby.
func murmurHash2_32(data []byte, seed uint32) (h uint32) {
	length := len(data)

	// Initialize the hash to a 'random' value
	h = seed ^ uint32(length)

	// Mix 4 bytes at a time into the hash
	var k uint32
	for length >= 4 {
		k = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
		h, k = mmix(h, k)

		data = data[4:]
		length -= 4
	}

	// Handle the last few bytes of the input array
	switch length {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(data[0])
		h *= M
	}

	// Do a few final mixes of the hash to ensure the last few bytes are well incorporated
	h ^= h >> 13
	h *= M
	h ^= h >> 15

	return
}
