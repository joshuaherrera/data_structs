package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

//need to be able to addnode, delete node, add at beginning, delete at beginnign
//delete some node, add after some node
//all linkedlist functions should be called from head node
/*func (head *Node) CreateNode() {
	head.Data = 0
	head.Next = nil
}*/

//add the value to the end of a linked list
func (list *LinkedList) Add(v int) {
	//create a temp node for the new value and a placeholder node
	temp, p := new(Node), new(Node)
	temp.Data = v
	//list is empty
	if list.Head.Next == nil && list.Head.Data == 0 {
		fmt.Println("Empty head")
		list.Head = temp
	} else {
		p = list.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = temp
	}
	//return list
}

//add a value after some node
func (list *LinkedList) AddAfter(n *Node, v int) {
	p, newN := new(Node), new(Node)
	newN.Data = v
	if list.Head.Next == nil && list.Head.Data == 0 {
		fmt.Println("Node not found")
		return
	}
	p = list.Head
	for p.Next != nil {
		if p == n {
			newN.Next = n.Next
			n.Next = newN
		}
		p = p.Next
	}
	//return list
}

//add a value at the start of a linked list
func (list *LinkedList) AddStart(v int) {
	newNode := new(Node)
	newNode.Data = v
	newNode.Next = list.Head
	list.Head = newNode
}

//remove a node from the end of a linkedlist
func (list *LinkedList) Remove() {
	p, pn := new(Node), new(Node)
	p = list.Head
	pn = p.Next
	for pn.Next != nil {
		p = p.Next
		pn = pn.Next
	}
	//at this point, pn is last node and need to remove
	p.Next = nil
}

//remove after some node, assume node is in the list
func (list *LinkedList) RemoveAfter(n *Node) {
	obsolete := new(Node)
	obsolete = n.Next
	n.Next = n.Next.Next
	obsolete.Next = nil
}

//remove at beginning of linkedlist
func (list *LinkedList) RemoveStart() {
	obsolete := new(Node)
	obsolete = list.Head
	list.Head = list.Head.Next
	obsolete.Next = nil
}

//check if value in linked list
func (list *LinkedList) Contains(v int) bool {
	p := new(Node)
	p = list.Head
	for p != nil {
		if p.Data == v {
			return true
		}
		p = p.Next
	}
	return false
}

//return some value
func (list *LinkedList) Get(v int) (*Node, bool) {
	p := new(Node)
	p = list.Head
	for p != nil {
		if p.Data == v {
			return p, true
		}
		p = p.Next
	}
	return nil, false
}

//traverse list and print values
func (list *LinkedList) Traverse() {
	p := new(Node)
	p = list.Head
	//what if last node has var data containing 0?
	if p.Data == 0 && p.Next == nil {
		fmt.Println("Linked List is empty")
		return
	}
	for p != nil {
		fmt.Printf("%v->", p.Data)
		p = p.Next
	}
	fmt.Println("|")
}
