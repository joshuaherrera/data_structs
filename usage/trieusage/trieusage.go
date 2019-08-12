package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/trie"
)

func main() {
	fmt.Println("testing trie")

	t := new(trie.Trie)
	t.Init()
	fmt.Println(t.Root())
	t.Insert("abc", 3)
	t.Insert("joshua", 7)
	t.Insert("shell", 9)
	t.Insert("shore", 19)
	fmt.Println("searching for shell's value: ", t.Find("shell"))
	fmt.Println("searching for shores's value: ", t.Find("shore"))
	fmt.Println(t.Children())
	t.Remove(t.Root(), "shell", 0)
	fmt.Println("Searching for shell's val: ", t.Find("shell"))
}
