package redblacktree

import (
	"fmt"
)

type color_t int

const (
	BLACK color_t = iota //iota starts at 0 for enum
	RED
)

type Node struct {
	/*Since the data in this implementation are integers,
	  we can use them as the keys. Otherwise, we'd have a
	  new int variable, key, that keeps sorted order, along
	  with the value to store.
	*/
	color  color_t
	data   int
	left   *Node
	right  *Node
	parent *Node
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
	root *Node
	size int
}

func (tree *RedBlackTree) GetParent(n *Node) *Node {
	return n.parent
}

func (tree *RedBlackTree) GetGrandParent(n *Node) *Node {
	p := tree.GetParent(n)
	if p == nil {
		return nil
	}
	return tree.GetParent(p)
}

func (tree *RedBlackTree) GetSibling(n *Node) *Node {
	p := tree.GetParent(n)
	if p == nil {
		//at root
		return nil
	}
	if n == p.left {
		return p.right
	} else {
		return p.left
	}
}

func (tree *RedBlackTree) GetUncle(n *Node) *Node {
	p := tree.GetParent(n)
	g := tree.GetGrandParent(n)
	if g == nil {
		return nil
	} else {
		return tree.GetSibling(p)
	}
}

func (tree *RedBlackTree) RotateLeft(n *Node) {
	fmt.Println("Rotating Left")
	newNode := n.right
	p := tree.GetParent(n)
	if newNode == nil {
		panic("leaves cannot be internal nodes due to emptiness")
	}
	n.right = newNode.left
	newNode.left = n
	n.parent = newNode
	if n.right != nil {
		n.right.parent = n
	}
	//check if n is root
	if p != nil {
		if n == p.left {
			p.left = newNode
		} else if n == p.right {
			p.right = newNode
		}
	}
	newNode.parent = p
}

func (tree *RedBlackTree) RotateRight(n *Node) {
	fmt.Println("Rotating Right")
	newNode := n.left
	p := tree.GetParent(n)
	if newNode == nil {
		panic("leaves cannot be internal nodes due to emptiness")
	}
	n.left = newNode.right
	newNode.right = n
	n.parent = newNode

	if n.left != nil {
		n.left.parent = n
	}

	if p != nil {
		if n == p.left {
			p.left = newNode
		} else if n == p.right {
			p.right = newNode
		}
	}
	newNode.parent = p
}

func (tree *RedBlackTree) InOrder(n *Node) {
	//fmt.Println("Calling InOrder traversal")
	if n == nil {
		return
	}
	tree.InOrder(n.left)
	fmt.Printf("Value: %v, color: %v ", n.data, n.color)
	tree.InOrder(n.right)
}

func (tree *RedBlackTree) PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("Value: %v , color: %v ", n.data, n.color)
	tree.PreOrder(n.left)
	tree.PreOrder(n.right)
}

func (tree *RedBlackTree) PostOrder(n *Node) {
	if n == nil {
		return
	}
	tree.PostOrder(n.left)
	tree.PostOrder(n.right)
	fmt.Printf("Value: %v , color: %v ", n.data, n.color)

}

func (tree *RedBlackTree) Search(val int) (*Node, bool) {
	fmt.Println("\nInitiating search")
	if tree.root == nil {
		return nil, false
	}
	currNode := tree.root

	for currNode != nil {
		if val == currNode.data {
			return currNode, true
		} else if val < currNode.data {
			currNode = currNode.left
		} else if val > currNode.data {
			currNode = currNode.right
		}
	}
	return nil, false

}

func (tree *RedBlackTree) Insert(val int) {
	fmt.Println("\nTesting Insertion")
	n := &Node{data: val}
	//insert node into tree
	root := tree.root
	if root == nil {
		n.parent = nil
		n.left = nil
		n.right = nil
		n.color = RED
		tree.root = n
	} else {
		tree.InsertRecurse(root, n)
	}
	//repair any Red-Black violations
	tree.InsertRepair(n)
	//return root to work with cases 4a/4b

	root = n
	for tree.GetParent(root) != nil {
		root = tree.GetParent(root)
	}
	tree.root = root
	tree.size++

}

func (tree *RedBlackTree) InsertRecurse(root, n *Node) {
	fmt.Println("Calling InsertRecurse")
	//recursively descend tree to find leaf node.
	if root != nil && n.data < root.data {
		if root.left != nil {
			//fmt.Println("\nGoing down left branch")
			tree.InsertRecurse(root.left, n)
			return
		} else {
			fmt.Println("Adding to left branch")
			root.left = n
		}
	} else if root != nil {
		if root.right != nil {
			//fmt.Println("\nGoing down right branch")
			tree.InsertRecurse(root.right, n)
			return
		} else {
			fmt.Println("Adding to right branch")
			root.right = n
		}
	}
	//set node values
	n.parent = root
	n.left = nil
	n.right = nil
	n.color = RED

}

func (tree *RedBlackTree) InsertRepair(n *Node) {
	fmt.Println("Calling InsertRepair")
	//violate any RB violations starting from newly added node
	if tree.GetParent(n) == nil {
		tree.InsertCase1(n)
	} else if tree.GetParent(n).color == BLACK {
		tree.InsertCase2(n)
	} else if tree.GetUncle(n) != nil && tree.GetUncle(n).color == RED {
		tree.InsertCase3(n)
	} else {
		tree.InsertCase4a(n)
	}
}

func (tree *RedBlackTree) InsertCase1(n *Node) {
	fmt.Println("Calling InsertCase1")
	//case 1: n is root and is currently RED
	if tree.GetParent(n) == nil {
		n.color = BLACK
	}
}

func (tree *RedBlackTree) InsertCase2(n *Node) {
	fmt.Println("Calling InsertCase2")
	//tree is valid do nothing
	return
}

func (tree *RedBlackTree) InsertCase3(n *Node) {
	fmt.Println("Calling InsertCase3")
	//case 3: parent and uncle are red
	//make them black and repaint grandparent red.
	//may need to repair grandparent if it's a root,
	//or for property 4 where children of red node must be black
	tree.GetParent(n).color = BLACK
	tree.GetUncle(n).color = BLACK
	tree.GetGrandParent(n).color = RED
	tree.InsertRepair(tree.GetGrandParent(n))
}

func (tree *RedBlackTree) InsertCase4a(n *Node) {
	fmt.Println("Calling InsertCase4a")
	//case 4a: parent is red, but uncle is black.
	//must put node n into grandparent position. wont
	//work if n is on inside of the tree. must do a rotation
	//then move to case 4b
	p := tree.GetParent(n)
	g := tree.GetGrandParent(n)

	if n == p.right && p == g.left {
		tree.RotateLeft(p)
		n = n.left
	} else if n == p.left && p == g.right {
		tree.RotateRight(p)
		n = n.right
	}
	tree.InsertCase4b(n)
}

func (tree *RedBlackTree) InsertCase4b(n *Node) {
	fmt.Println("Calling InsertCase4b")
	//case 4b: n is now on outside of tree. perform rt rotation
	//on grandparent. p is not parent of n and former g. g is black.
	//swap colors with p and g
	p := tree.GetParent(n)
	g := tree.GetGrandParent(n)

	if n == p.left {
		tree.RotateRight(g)
	} else {
		tree.RotateLeft(g)
	}
	p.color = BLACK
	g.color = RED
}

func (tree *RedBlackTree) Clear() {
	tree.root = nil
	tree.size = 0
}

func (tree *RedBlackTree) FindMin(n *Node) *Node {
	fmt.Println("Testing FindMin")
	/*return min node. used for getting left-most child
	  of right subtree
	*/
	currNode := n
	for currNode.left != nil {
		currNode = currNode.left
	}
	return currNode
}

func (tree *RedBlackTree) ReplaceNode(n, child *Node) {
	fmt.Println("Testing ReplaceNode")
	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = child
		} else {
			n.parent.right = child
		}
	}
	if child != nil {
		child.parent = n.parent
	}
}

func (tree *RedBlackTree) Delete(val int) {
	fmt.Println("\nTesting deletion")
	if tree.root == nil {
		return
	}

	//find the node to delete
	currNode, boolean := tree.Search(val)
	fmt.Println("Value found? ", boolean)
	//check if has 2 children, and address if so
	if currNode.left != nil && currNode.right != nil {
		successor := tree.FindMin(currNode.right)
		currNode.data = successor.data
		//change currNode to continue with next cases
		currNode = successor
	}
	//deal with case where 0 or 1 children, returns new root
	// could we just do tree.root = currNode.DeleteOneChild?
	tree.root = tree.DeleteOneChild(currNode)
	tree.size--
	//figure out if need to modify to work with my code
	//may need to bubble up from tree.root to find root.

}

func (tree *RedBlackTree) DeleteOneChild(n *Node) *Node {
	fmt.Println("Testing DeleteOneChild")
	var child, cparent *Node
	if n.right == nil {
		child = n.left
	} else {
		child = n.right
	}
	cparent = n.parent
	tree.ReplaceNode(n, child)
	fmt.Println("Exitted ReplaceNode")
	if n.color == BLACK {
		if GetColor(child) == RED {
			child.color = BLACK
		} else {
			tree.DeleteCase1(child, cparent)
		}
	}
	//search for root and return
	//currently child is null. need a node still on tree to get root
	var root *Node
	if child == nil {
		//n's parent should be valid node unless is root
		root = tree.GetParent(n)
	} else {
		root = child
	}
	//might be able to get rid of conditionals above by checking parent
	//ie above can set root to cparent
	for tree.GetParent(root) != nil {
		root = tree.GetParent(root)
	}
	//'delete' n and return the new root
	n = nil
	return root

}

func (tree *RedBlackTree) DeleteCase1(n, parent *Node) {
	fmt.Println("Testing DeleteCase1")
	//terminal case...
	//to avoid nil memorys, maybe i could check if n is null and go to case 2.
	//but that just passes problem along
	//I could pass the child's new parent instead.. or keep track of it
	//and use it for the cases instead of the nil value
	if parent != nil {
		tree.DeleteCase2(n, parent)
	}
}

func (tree *RedBlackTree) DeleteCase2(n, parent *Node) {
	fmt.Println("Testing DeleteCase2")
	//could check if n nil, if so get parents child for s
	var s *Node
	if n != nil {
		s = tree.GetSibling(n)
	} else if n == parent.right {
		s = parent.left
	} else {
		s = parent.right
	}
	if GetColor(s) == RED {
		parent.color = RED
		s.color = BLACK
		if n == parent.left {
			tree.RotateLeft(parent)
		} else {
			tree.RotateRight(parent)
		}
	}
	tree.DeleteCase3(n, parent)
}

func (tree *RedBlackTree) DeleteCase3(n, parent *Node) {
	fmt.Println("Testing DeleteCase3")
	var s *Node
	if n != nil {
		s = tree.GetSibling(n)
	} else if n == parent.right {
		s = parent.left
	} else {
		s = parent.right
	}
	if GetColor(parent) == BLACK && GetColor(s) == BLACK &&
		GetColor(s.left) == BLACK && GetColor(s.right) == BLACK {
		s.color = RED
		tree.DeleteCase1(parent, parent.parent)
	} else {
		tree.DeleteCase4(n, parent)
	}
}

func (tree *RedBlackTree) DeleteCase4(n, parent *Node) {
	fmt.Println("Testing DeleteCase4")
	//terminal case
	var s *Node
	if n != nil {
		s = tree.GetSibling(n)
	} else if n == parent.right {
		s = parent.left
	} else {
		s = parent.right
	}

	if GetColor(parent) == RED && GetColor(s) == BLACK &&
		GetColor(s.left) == BLACK && GetColor(s.right) == BLACK {
		s.color = RED
		parent.color = BLACK
	} else {
		tree.DeleteCase5(n, parent)
	}
}

func (tree *RedBlackTree) DeleteCase5(n, parent *Node) {
	fmt.Println("Testing DeleteCase5")
	var s *Node
	if n != nil {
		s = tree.GetSibling(n)
	} else if n == parent.right {
		s = parent.left
	} else {
		s = parent.right
	}
	if GetColor(s) == BLACK {
		if n == parent.left && GetColor(s.right) == BLACK &&
			GetColor(s.left) == RED {
			s.color = RED
			s.left.color = BLACK
			tree.RotateRight(s)
		} else if n == parent.right && GetColor(s.left) == BLACK &&
			GetColor(s.right) == RED {
			s.color = RED
			s.right.color = BLACK
			tree.RotateLeft(s)
		}
		tree.DeleteCase6(n, parent)
	}
}

func (tree *RedBlackTree) DeleteCase6(n, parent *Node) {
	fmt.Println("Testing DeleteCase6")
	//terminal case
	var s *Node
	if n != nil {
		s = tree.GetSibling(n)
	} else if n == parent.right {
		s = parent.left
	} else {
		s = parent.right
	}
	s.color = parent.color
	parent.color = BLACK

	if n == parent.left {
		s.right.color = BLACK
		tree.RotateLeft(parent)
	} else {
		s.left.color = BLACK
		tree.RotateRight(parent)
	}

}

func (tree *RedBlackTree) GetRoot() *Node {
	return tree.root
}

func (tree *RedBlackTree) GetSize() int {
	return tree.size
}

func GetColor(n *Node) color_t {
	if n == nil {
		return BLACK
	}
	return n.color
}
