package bst

import (
	"fmt"
)

type Node struct {
	/*Since the Data in this implementation are integers,
	  we can use them as the keys. Otherwise, we'd have a
	  new int variable, key, that keeps sorted order, along
	  with the value to store.
	*/
	Data   int
	Left   *Node
	Right  *Node
	Parent *Node
}

type Bst struct {
	Root *Node
}

func (bst *Bst) Init(val int) {
	fmt.Println("Init Root")
	n := new(Node)
	n.CreateNode(val)
	n.Parent = nil
	bst.Root = n
}

func (n *Node) CreateNode(v int) {
	/*looks unneeded since can do &Node*/
	n.Data = v
	n.Left = nil
	n.Right = nil
}

func (tree *Bst) Insert(val int) {
	fmt.Println("Calling insert")
	if tree.Root == nil {
		tree.Root = &Node{Data: val, Parent: nil}
		return
	}

	currNode := tree.Root

	for {
		if val < currNode.Data {
			if currNode.Left == nil {
				currNode.Left = &Node{Data: val, Parent: currNode}
				return
			}
			currNode = currNode.Left
		} else {
			if currNode.Right == nil {
				currNode.Right = &Node{Data: val, Parent: currNode}
				return
			}
			currNode = currNode.Right
		}
	}

}

func (tree *Bst) Search(val int) (*Node, bool) {
	fmt.Println("Initiating search")
	if tree.Root == nil {
		return nil, false
	}
	currNode := tree.Root

	for currNode != nil {
		if val == currNode.Data {
			return currNode, true
		} else if val < currNode.Data {
			currNode = currNode.Left
		} else if val > currNode.Data {
			currNode = currNode.Right
		}
	}
	return nil, false

}

func (n *Node) FindMin() *Node {
	fmt.Println("Testing FindMin")
	/*return min node. used for getting left-most child
	  of right subtree
	*/
	currNode := n
	for currNode.Left != nil {
		currNode = currNode.Left
	}
	return currNode
}

func (n *Node) ReplaceNodeInParent(newVal *Node) {
	fmt.Println("Testing ReplaceNodeInParent")
	if n.Parent != nil {
		if n == n.Parent.Left {
			n.Parent.Left = newVal
		} else {
			n.Parent.Right = newVal
		}
	}
	if newVal != nil {
		newVal.Parent = n.Parent
	}
}

func (tree *Bst) Delete(val int) {
	fmt.Println("\nTesting deletion")
	if tree.Root == nil {
		return
	}

	//find the node to delete
	currNode, boolean := tree.Search(val)
	fmt.Println("Value found? ", boolean)
	//check if has 2 children, and address if so
	if currNode.Left != nil && currNode.Right != nil {
		successor := currNode.Right.FindMin()
		currNode.Data = successor.Data
		//change currNode to continue with next cases
		currNode = successor
	}
	//deal with case where 0 or 1 children
	if currNode.Left != nil {
		currNode.ReplaceNodeInParent(currNode.Left)
	} else if currNode.Right != nil {
		currNode.ReplaceNodeInParent(currNode.Right)
	} else {
		currNode.ReplaceNodeInParent(nil)
	}

}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%v ", n.Data)
	n.Right.InOrder()
}
