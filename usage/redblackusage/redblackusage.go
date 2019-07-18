package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/redblacktree"
)

func main() {
	fmt.Println("Starting redblacktree testing")
	b := new(redblacktree.RedBlackTree)
	b.Insert(25)
	//b.Root.InOrder()
	b.Root.PreOrder()
	b.Insert(75)
	//b.Root.InOrder()
	b.Root.PreOrder()
	b.Insert(100)
	//b.Root.InOrder()
	b.Root.PreOrder()
	b.Insert(10)
	//b.Root.InOrder()
	b.Root.PreOrder()
	b.Insert(30)
	//b.Root.InOrder()
	b.Root.PreOrder()
	b.Insert(60)
	//b.Root.InOrder()
	b.Root.PreOrder()
	fmt.Println()
	b.Insert(70)
	//b.Root.InOrder()
	fmt.Println()
	b.Root.PreOrder()
	fmt.Println()
	b.Delete(70)
	b.Root.PreOrder()
	fmt.Println()

	node, bool_val := b.Search(75)
	fmt.Println(node, bool_val)
}
