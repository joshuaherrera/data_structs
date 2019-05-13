package binarytree

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

func (bt *BinaryTree) Init() {
	bt.data = 0
	bt.left = nil
	bt.right = nil
}

func (bt *BinaryTree) CreateNode(v int) {
	bt.data = v
	bt.left = nil
	bt.right = nil
}

func (bt *BinaryTree) InOrder() {
	if bt == nil {
		return
	}
	bt.left.InOrder()
	fmt.Printf("%v ", bt.data)
	bt.right.InOrder()
}

func (bt *BinaryTree) PreOrder() {
	if bt == nil {
		return
	}
	fmt.Printf("%v ", bt.data)
	bt.left.PreOrder()
	bt.right.PreOrder()
}

func (bt *BinaryTree) PostOrder() {
	if bt == nil {
		return
	}
	bt.left.PostOrder()
	bt.right.PostOrder()
	fmt.Printf("%v ", bt.data)
}

//TODO: deletion

func (bt *BinaryTree) Insert(v int) {
	var q []*BinaryTree
	q = append(q, bt)

	for len(q) > 0 {
		temp := new(BinaryTree)
		temp = q[0]
		q[0] = nil
		q = q[1:]
		//check left and right
		if temp.left == nil {
			newNode := new(BinaryTree)
			newNode.CreateNode(v)
			temp.left = newNode
			break
		} else {
			q = append(q, temp.left)
		}
		if temp.right == nil {
			newNode := new(BinaryTree)
			newNode.CreateNode(v)
			temp.right = newNode
			break
		} else {
			q = append(q, temp.right)
		}
	}
}
