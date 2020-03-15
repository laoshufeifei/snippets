package main

import (
	"fmt"
	"gossips/arraylist"
)

func main() {
	fmt.Println("hello gossips")

	l := arraylist.New("a", "b", "c")
	fmt.Println(l.Size())
}
