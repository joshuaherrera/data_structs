package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/hashtable"
)

func main() {
	ht := new(hashtable.HashTable)
	ht.Init()
	fmt.Println("Testing Put")
	ht.Put("a", 1)
	ht.Put("b", 3)
	ht.Put("c", 131)
	ht.Put("d", 121)
	ht.Put("e", 1323)
	ht.Put("f", 133)
	ht.PrintBuckets()
	fmt.Println("Testing Get")
	fmt.Println(ht.Get("d"))
	fmt.Println(ht.Get("a"))
	fmt.Println(ht.Get("z"))
	fmt.Println("Testing Remove")
	ht.Remove("d")
	fmt.Println(ht.Get("d"))
	ht.Remove("b")
	fmt.Println(ht.Get("b"))
	ht.PrintBuckets()
}
