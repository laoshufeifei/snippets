package main

import (
	"fmt"
	"gossips/src/mysort"
)

func updateInts(s []int) {
	s[0] = 100
	fmt.Printf("2: %p, len: %d, cap: %d\n", &s, len(s), cap(s))

	s = append(s, 666)
	fmt.Printf("3: %p, len: %d, cap: %d\n", &s, len(s), cap(s))
}

func testSlice() {
	// s := []int{1, 2, 3}
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3
	fmt.Printf("1: %p, len: %d, cap: %d\n", &s, len(s), cap(s))

	updateInts(s)
	fmt.Printf("4: %p, len: %d, cap: %d\n", &s, len(s), cap(s))

	fmt.Println(s)
}

func main() {
	mysort.BenchmarkTestSorters()
}
