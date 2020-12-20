package binaryindexedtree

func lowbit(x int) int {
	return x & -x
}
