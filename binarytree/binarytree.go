package binarytree

import (
	"fmt"
)

type BinaryTree struct {
	data  interface{}
	left  *BinaryTree
	right *BinaryTree
}

func (bt *BinaryTree) Init() {
	bt.data = 0
	bt.left = nil
	bt.right = nil
}

func (bt *BinaryTree) CreateNode(v interface{}) {
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

func (bt *BinaryTree) LevelOrder() {
	if bt == nil {
		return
	}
	q := make([]*BinaryTree, 0)
	discovered := make(map[*BinaryTree]bool)
	discovered[bt] = true
	q = append(q, bt)
	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		fmt.Printf("%v ", temp.data)
		if _, ok := discovered[temp.left]; !ok && temp.left != nil {
			q = append(q, temp.left)
		}
		if _, ok := discovered[temp.right]; !ok && temp.right != nil {
			q = append(q, temp.right)
		}

	}

}

//TODO: deletion

func (bt *BinaryTree) Insert(v interface{}) {
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
