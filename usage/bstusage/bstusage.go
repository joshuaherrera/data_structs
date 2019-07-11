package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/bst"
)

func main() {
	fmt.Println("Starting BST testing")
	b := new(bst.Bst)
	b.Init(50)
	b.Insert(25)
	b.Insert(75)
	b.Insert(100)
	b.Insert(10)
	b.Insert(30)
	b.Insert(60)
	b.Root.InOrder()
	/*	node := new(bst.Node)
		var bool_val bool*/
	node, bool_val := b.Search(30)
	fmt.Println(node, bool_val)
	b.Delete(75)
	b.Root.InOrder()
	b.Delete(25)
	b.Root.InOrder()
	b.Delete(50)
	b.Root.InOrder()

}
