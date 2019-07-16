package redblacktree

import (
	"fmt"
)

type Color_t int

const (
	BLACK Color_t = iota //iota starts at 0 for enum
	RED
)

type Node struct {
	/*Since the Data in this implementation are integers,
	  we can use them as the keys. Otherwise, we'd have a
	  new int variable, key, that keeps sorted order, along
	  with the value to store.
	*/
	Color  Color_t
	Data   int
	Left   *Node
	Right  *Node
	Parent *Node
}

type RedBlackTree struct {
	/*	Properties of a Red-Black Tree (source: wikipedia):
		1. Each node is either red or black.
		2. The root is black. This rule is sometimes omitted.
		   Since the root can always be changed from red to black,
		   but not necessarily vice versa, this rule has little effect on analysis.
		3. All leaves (NIL) are black.
		4. If a node is red, then both its children are black.
		5. Every path from a given node to any of its descendant NIL nodes
		   contains the same number of black nodes.
	*/
	Root *Node
}

func (n *Node) GetParent() *Node {
	return n.Parent
}

func (n *Node) GetGrandParent() *Node {
	p := n.GetParent()
	if p == nil {
		return nil
	}
	return p.GetParent()
}

func (n *Node) GetSibling() *Node {
	p := n.GetParent()
	if p == nil {
		//at root
		return nil
	}
	if n == p.Left {
		return p.Right
	} else {
		return p.Left
	}
}

func (n *Node) GetUncle() *Node {
	p := n.GetParent()
	g := n.GetGrandParent()
	if g == nil {
		return nil
	} else {
		return p.GetSibling()
	}
}

func (n *Node) RotateLeft() {
	newNode := n.Right
	p := n.GetParent()
	if newNode == nil {
		panic("leaves cannot be internal nodes due to emptiness")
	}
	n.Right = newNode.Left
	newNode.Left = n
	n.Parent = newNode
	if n.Right != nil {
		n.Right.Parent = n
	}
	//check if n is root
	if p != nil {
		if n == p.Left {
			p.Left = newNode
		} else if n == p.Right {
			p.Right = newNode
		}
	}
	newNode.Parent = p
}

func (n *Node) RotateRight() {
	newNode := n.Left
	p := n.GetParent()
	if newNode == nil {
		panic("leaves cannot be internal nodes due to emptiness")
	}
	n.Left = newNode.Right
	newNode.Right = n
	n.Parent = newNode

	if n.Left != nil {
		n.Left.Parent = n
	}

	if p != nil {
		if n == p.Left {
			p.Left = newNode
		} else if n == p.Right {
			p.Right = newNode
		}
	}
	newNode.Parent = p
}

func (n *Node) InOrder() {
	//fmt.Println("Calling InOrder traversal")
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("Value: %v, Color: %v ", n.Data, n.Color)
	n.Right.InOrder()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("Value: %v , Color: %v ", n.Data, n.Color)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("Value: %v , Color: %v ", n.Data, n.Color)

}

func (tree *RedBlackTree) Search(val int) (*Node, bool) {
	fmt.Println("\nInitiating search")
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

func (tree *RedBlackTree) Insert(val int) *Node {
	fmt.Println("\nTesting Insertion")
	n := &Node{Data: val}
	//insert node into tree
	root := tree.Root
	if root == nil {
		n.Parent = nil
		n.Left = nil
		n.Right = nil
		n.Color = RED
		tree.Root = n
	} else {
		root.InsertRecurse(n)
	}
	//repair any Red-Black violations
	n.InsertRepair()
	//return root to work with cases 4a/4b

	root = n
	for root.GetParent() != nil {
		root = root.GetParent()
	}
	return root

}

func (root *Node) InsertRecurse(n *Node) {
	fmt.Println("Calling InsertRecurse")
	//recursively descend tree to find leaf node.
	if root != nil && n.Data < root.Data {
		if root.Left != nil {
			fmt.Println("\nGoing down left branch")
			root.Left.InsertRecurse(n)
			return
		} else {
			fmt.Println("Adding to left branch")
			root.Left = n
		}
	} else if root != nil {
		if root.Right != nil {
			fmt.Println("\nGoing down right branch")
			root.Right.InsertRecurse(n)
			return
		} else {
			fmt.Println("Adding to right branch")
			root.Right = n
		}
	}
	//set node values
	n.Parent = root
	n.Left = nil
	n.Right = nil
	n.Color = RED

}

func (n *Node) InsertRepair() {
	fmt.Println("Calling InsertRepair")
	//violate any RB violations starting from newly added node
	if n.GetParent() == nil {
		n.InsertCase1()
	} else if n.GetParent().Color == BLACK {
		n.InsertCase2()
	} else if n.GetUncle() != nil && n.GetUncle().Color == RED {
		n.InsertCase3()
	} else {
		n.InsertCase4a()
	}
}

func (n *Node) InsertCase1() {
	fmt.Println("Calling InsertCase1")
	//case 1: n is root and is currently RED
	if n.GetParent() == nil {
		n.Color = BLACK
	}
}

func (n *Node) InsertCase2() {
	fmt.Println("Calling InsertCase2")
	//tree is valid do nothing
	return
}

func (n *Node) InsertCase3() {
	fmt.Println("Calling InsertCase3")
	//case 3: parent and uncle are red
	//make them black and repaint grandparent red.
	//may need to repair grandparent if it's a root,
	//or for property 4 where children of red node must be black
	n.GetParent().Color = BLACK
	n.GetUncle().Color = BLACK
	n.GetGrandParent().Color = RED
	n.GetGrandParent().InsertRepair()
}

func (n *Node) InsertCase4a() {
	fmt.Println("Calling InsertCase4a")
	//case 4a: parent is red, but uncle is black.
	//must put node n into grandparent position. wont
	//work if n is on inside of the tree. must do a rotation
	//then move to case 4b
	p := n.GetParent()
	g := n.GetGrandParent()

	if n == p.Right && p == g.Left {
		p.RotateLeft()
		n = n.Left
	} else if n == p.Left && p == g.Right {
		p.RotateRight()
		n = n.Right
	}
	n.InsertCase4b()
}

func (n *Node) InsertCase4b() {
	fmt.Println("Calling InsertCase4b")
	//case 4b: n is now on outside of tree. perform rt rotation
	//on grandparent. p is not parent of n and former g. g is black.
	//swap colors with p and g
	p := n.GetParent()
	g := n.GetGrandParent()

	if n == p.Left {
		g.RotateRight()
	} else {
		g.RotateLeft()
	}
	p.Color = BLACK
	g.Color = RED
}
