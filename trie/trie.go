package trie

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	children map[string]*Node
	value    interface{}
}

func (t *Trie) Init() {
	t.root = newNode()
}

func (t *Trie) Find(key string) interface{} {
	fmt.Println("testing find")
	node := t.root
	for _, r := range key {
		char := string(r) //convert rune value to str
		if _, ok := node.children[char]; ok {
			fmt.Println("found node")
			node = node.children[char]
		} else {
			return nil
		}
	}
	return node.value
}

func (t *Trie) Insert(key string, v interface{}) {
	fmt.Println("testing insert")
	node := t.root
	for _, r := range key {
		char := string(r)
		//fmt.Println(r, char)
		if _, ok := node.children[char]; ok == false {
			node.children[char] = newNode()
		}
		node = node.children[char]
	}
	node.value = v
}

func (t *Trie) isEmpty(n *Node) bool {
	if len(n.children) == 0 {
		return true
	}
	return false
}

func (t *Trie) Remove(node *Node, key string, depth int) *Node {
	fmt.Println("testing remove")
	if node == nil {
		return nil
	}

	//check if were at the last char of key
	if depth == len(key) {
		//values only found at end of keys, must remove
		if node.value != nil {
			node.value = nil
		}

		//if node has no children that make up keys
		if t.isEmpty(node) {
			//fmt.Println("at last char, no children")
			node.children = nil
			node = nil
		}
		return node
	}

	//recurse using each character in the key
	charKey := string([]rune(key)[depth])
	node.children[charKey] = t.Remove(node.children[charKey], key, depth+1)

	if node.children[charKey] == nil && node.value == nil {
		//fmt.Println("Found preceding char and no val")
		node.children = nil
		node = nil
	}

	return node
}

func (t *Trie) Children() map[string]*Node {
	return t.root.children
}

func (t *Trie) Root() *Node {
	return t.root
}

func newNode() *Node {
	return &Node{
		children: make(map[string]*Node),
	}
}
