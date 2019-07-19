package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/redblacktree"
)

func main() {
	fmt.Println("Starting redblacktree testing")
	tree := new(redblacktree.RedBlackTree)
	tree.Insert(25)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(75)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(100)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(10)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(30)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(60)
	//tree.Root.InOrder()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(70)
	fmt.Println("\nSize is: ", tree.GetSize())
	//tree.Root.InOrder()
	fmt.Println()
	tree.PreOrder(tree.GetRoot())
	tree.Insert(150)
	fmt.Println("\nSize is: ", tree.GetSize())
	//tree.Root.InOrder()
	fmt.Println()
	tree.PreOrder(tree.GetRoot())
	fmt.Println()
	tree.Delete(70)
	tree.PreOrder(tree.GetRoot())
	fmt.Println("\nSize is: ", tree.GetSize())
	tree.Delete(100)
	tree.PreOrder(tree.GetRoot())
	fmt.Println("\nSize is: ", tree.GetSize())

	/*	node, bool_val := tree.Search(75)
		fmt.Println(node, bool_val)
	*/
}
