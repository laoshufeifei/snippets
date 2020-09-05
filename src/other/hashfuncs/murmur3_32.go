package hashfuncs

// https://sites.google.com/site/murmurhash/
// https://github.com/spaolacci/murmur3/blob/master/murmur32.go

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
)

func murmurHash3_32(data []byte, seed uint32) uint32 {
	length := uint32(len(data))
	oldlen := length
	h := seed

	var k uint32
	for length >= 4 {
		k1 := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24

		k1 *= c1
		k1 = k1<<15 | k1>>17
		k1 *= c2

		h ^= k1
		h = h<<13 | h>>19
		h = h*4 + h + 0xe6546b64

		data = data[4:]
		length -= 4
	}

	k = 0
	switch length {
	case 3:
		k ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		k ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		k ^= uint32(data[0])
		k *= c1
		k = k<<15 | k>>17
		k *= c2
		h ^= k
	}

	h ^= oldlen

	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}
