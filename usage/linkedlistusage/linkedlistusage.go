package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/linkedlist"
)

func main() {
	l := new(linkedlist.LinkedList)
	l.Head = new(linkedlist.Node)
	//fmt.Printf("%v of type %T\n", l.Head, l.Head)
	l.Add(5)
	//fmt.Printf("%v of type %T\n", l.Head, l.Head)
	l.Add(10)
	//fmt.Printf("%v of type %T\n", l.Head.Next, l.Head.Next)
	if l.Contains(5) {
		fmt.Println("Contains 5")
	}
	l.Traverse()
	l.AddStart(7)
	l.Traverse()
	p, b := l.Get(5)
	if b {
		l.AddAfter(p, 99)
	}
	l.Traverse()
	l.RemoveAfter(p)
	l.Traverse()
	l.Add(9999)
	l.Traverse()
	l.Remove()
	l.Traverse()
	l.RemoveStart()
	l.Traverse()
}
