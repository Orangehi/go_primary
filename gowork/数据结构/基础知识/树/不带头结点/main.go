package main

import (
	"container/list"
	"fmt"
)

//一条 双向循环链表
type Lister struct {
	L *list.List
}

//rewrting String()string
func (l *Lister) String() string {

	length := l.L.Len()
	fmt.Println(length)
	h := l.L.Front()
	h1 := h.Value.(*BinaryTreeNode)
	st := fmt.Sprintf("the list-length:%d\r\n", length)
	for j := 0; j < length; j++ {

		h1 = h.Value.(*BinaryTreeNode)
		st += h1.date
		st += "\r\n"
		h = h.Next()
	}
	return st

}

//binarytreenode
type BinaryTreeNode struct {
	date   string
	parent *BinaryTreeNode //parent node
	lchild *BinaryTreeNode //left child node
	rchild *BinaryTreeNode // right child node
	height int             //以该节点为根的高度
	size   int             //该节点的子孙数(包括该节点)
}

//construct function
func MakeTree() *BinaryTreeNode {

	n := &BinaryTreeNode{date: "this is root", height: 1, size: 1}
	n.parent = n
	return n
}

//set lchild node
func (n *BinaryTreeNode) SetLchild(l *BinaryTreeNode) bool {

	if n.lchild != nil {

		return false
	}

	n.lchild = l
	return true

}

//set rchild node
func (n *BinaryTreeNode) SetRchild(r *BinaryTreeNode) bool {

	if n.rchild != nil {

		return false
	}

	n.rchild = r
	return true

}

//二叉树的后序排列，左右中
func (t *BinaryTreeNode) TailOrder() *list.List {

	traversal := list.New()
	tailOrder(t, traversal)
	return traversal

}

func tailOrder(n *BinaryTreeNode, l *list.List) {

	if n == nil {

		return

	}

	tailOrder(n.lchild, l)
	tailOrder(n.rchild, l)
	l.PushBack(n)

}

//二叉树的前序排列 中左右
func (t *BinaryTreeNode) PrevOrder() *list.List {

	traversal := list.New()

	prevOrder(t, traversal)

	return traversal
}

//二叉树前序排列的递归函数
func prevOrder(n *BinaryTreeNode, l *list.List) {

	if n == nil {

		return

	}

	l.PushBack(n)
	prevOrder(n.lchild, l)
	prevOrder(n.rchild, l)

}

//二叉树的中序排列 （左中右） 存放到一个双向循环链表中
func (t *BinaryTreeNode) InOrder() *list.List {

	traversal := list.New()

	inOrder(t, traversal)

	return traversal

}

//二叉树的中序排列
func inOrder(n *BinaryTreeNode, l *list.List) {

	if n == nil {

		return

	}

	inOrder(n.lchild, l)
	l.PushBack(n)
	inOrder(n.rchild, l)

}
func main() {

	t := MakeTree()
	fmt.Println(t.date)

	n2 := &BinaryTreeNode{date: "this is node2"}
	t.SetLchild(n2)
	n3 := &BinaryTreeNode{date: "this is node3"}
	t.SetRchild(n3)
	n4 := &BinaryTreeNode{date: "this is node4"}
	t.lchild.SetLchild(n4)
	n5 := &BinaryTreeNode{date: "this is node5"}
	t.lchild.SetRchild(n5)
	n6 := &BinaryTreeNode{date: "this is node6"}
	t.rchild.SetLchild(n6)
	n7 := &BinaryTreeNode{date: "this is node7"}
	t.rchild.SetRchild(n7)

	l := t.InOrder()
	l1 := &Lister{L: l}
	fmt.Println(l1)
}
