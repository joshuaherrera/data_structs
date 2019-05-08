package linkedlist

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

func (list *LinkedList) Add(v int) *LinkedList {
	//create a temp node for the new value and a placeholder node
	temp, p := new(Node)
	temp.Data = v
	//list is empty
	if list.Head == nil {
		list.Head = temp
	} else {
		p = list.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = temp
	}
	return list
}

func (head *Node) AddAfter(n *Node, v int) *Node {
	p, newN := new(Node)
	newN.Data = v
	if head == nil {
		return head
	}
	p = head
	for p.Next != nil {
		if p == n {
			newN.Next = n.Next
			n.Next = newN
		}
		p = p.Next
	}
	return head
}

func (head *Node) AddStart(list *LinkedList)
