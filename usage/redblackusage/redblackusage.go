package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/redblacktree"
)

func main() {
	fmt.Println("Starting redblacktree testing")
	b := new(redblacktree.RedBlackTree)
	b.Root = b.Insert(25)
	b.Root.InOrder()
	b.Root.PreOrder()
	b.Root = b.Insert(75)
	b.Root.InOrder()
	b.Root.PreOrder()
	b.Root = b.Insert(100)
	b.Root.InOrder()
	b.Root.PreOrder()
	b.Root = b.Insert(10)
	b.Root.InOrder()
	b.Root.PreOrder()
	b.Root = b.Insert(30)
	b.Root.InOrder()
	b.Root.PreOrder()
	b.Root = b.Insert(60)
	b.Root.InOrder()
	b.Root.PreOrder()
	fmt.Println()
	b.Root.PreOrder()
	fmt.Println()
	b.Root = b.Insert(70)
	b.Root.InOrder()
	fmt.Println()
	b.Root.PreOrder()
	fmt.Println()

	/*	node, bool_val := b.Search(75)
		node2, bool_val2 := b.Search(100)
		fmt.Println(node, bool_val, node2, bool_val2)*/
}
