package bloomfilter

import (
	"hash/fnv"
	"math"
)

// Bloomfilter 假设就以字符串来进行计算
type Bloomfilter struct {
	hashCount uint64 // hash function count
	bitLength uint64
	bitSet    *Bits
}

// New ...
func New(n uint64, p float64) *Bloomfilter {
	ln2 := math.Ln2
	bitLength := uint64(-1. * float64(n) * math.Log(p) / (ln2 * ln2))
	// hashCount := uint64(float64(bitLength) / float64(n) * ln2)
	hashCount := uint64(-1. * math.Log2(p))

	b := &Bloomfilter{
		hashCount: hashCount,
		bitSet:    newBitSet(bitLength),
		bitLength: bitLength,
	}

	return b
}

// Put ...
func (b *Bloomfilter) Put(v string) {
	hashCode := fnvHashString(v)
	hashCode2 := hashCode >> 32

	for i := uint64(1); i <= b.hashCount; i++ {
		combined := hashCode + (uint64(i) * hashCode2)

		index := combined % uint64(b.bitLength)
		b.bitSet.set(index)
	}
}

// Contains ...
func (b *Bloomfilter) Contains(v string) bool {
	hashCode := fnvHashString(v)
	hashCode2 := hashCode >> 32

	for i := uint64(1); i <= b.hashCount; i++ {
		combined := hashCode + (uint64(i) * hashCode2)

		index := combined % uint64(b.bitLength)
		if !b.bitSet.get(index) {
			return false
		}
	}

	return true
}

func fnvHashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
