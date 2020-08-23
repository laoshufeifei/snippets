package other

// https://golang.org/src/hash/fnv/fnv.go
// https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function

import "hash/fnv"

func golangFnv32(bs []byte) uint32 {
	h := fnv.New32()
	h.Write(bs)
	return h.Sum32()
}

func golangFnv32a(bs []byte) uint32 {
	h := fnv.New32a()
	h.Write(bs)
	return h.Sum32()
}

func fnv132(bytes []byte) uint32 {
	var hash uint32 = 0x811c9dc5
	for _, b := range bytes {
		hash *= 16777619
		hash ^= uint32(b)
	}

	return hash
}

func fnv1a32(bytes []byte) uint32 {
	var hash uint32 = 0x811c9dc5
	for _, b := range bytes {
		hash ^= uint32(b)
		hash *= 16777619
	}

	return hash
}
