package bloomfilter

import "math"

// Bloomfilter ...
type Bloomfilter struct {
	hashCount int   // hash function count
	bitSize   int64 // bits size
	bits      []int64
}

// New ...
func New(p float64, n int64) *Bloomfilter {
	b := &Bloomfilter{
		bitSize:   -1 * n * int64(math.Log(p)/math.Ln2/math.Ln2),
		hashCount: -1 * int(math.Log2(p)),
	}

	return b
}

// Put ...
func (b *Bloomfilter) Put(v string) bool {
	return true
}

// Contains ...
func (b *Bloomfilter) Contains(v string) bool {
	return true
}
