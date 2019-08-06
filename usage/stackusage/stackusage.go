package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/stack"
)

func main() {
	s := new(stack.Stack)
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
