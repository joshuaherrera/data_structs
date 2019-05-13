package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/binarytree"
)

func main() {
	bt := new(binarytree.BinaryTree)
	bt.CreateNode(5)
	bt.InOrder()
	fmt.Println()
	bt.Insert(2)
	bt.Insert(3)
	bt.Insert(4)
	bt.Insert(6)
	bt.InOrder()
	fmt.Println()
	bt.PreOrder()
	fmt.Println()
	bt.PostOrder()
	fmt.Println()
}
