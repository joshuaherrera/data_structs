package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/stack"
)

func main() {
	s := stack.Stack{Items: make([]int, 0, 10), Top: 0}
	s.Push(5)
	s.Push(10)
	s.Push(15)
	fmt.Println(s.Items)
	fmt.Println("Peeking")
	fmt.Println(s.Peek())
	fmt.Println("Popping")
	fmt.Println(s.Pop())
	fmt.Println(s.Items)
}
