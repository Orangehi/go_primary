package main

import (
	"container/list"
	"fmt"
	"os"
	//"reflect"
)

//binarytree
type BinaryTree struct {
	root   *BinaryTreeNode //根节点
	height int             //二叉树的高度
	size   int             //二叉树的节点数
}

//binarytreenode
type BinaryTreeNode struct {
	date   string
	parent *BinaryTreeNode //parent node
	lchild *BinaryTreeNode //left child node
	rchild *BinaryTreeNode // right child node
	//	height int             //以该节点为根的高度
	//	size   int             //该节点的子孙数(包括该节点)
}

//BinaryTree construct function
func MakeTree() *BinaryTree {
	n := &BinaryTreeNode{date: "this is root"}
	t := &BinaryTree{root: n, height: 1, size: 1}
	return t

}

//获取该节点的lchlid node
func GetLchildNode(n *BinaryTreeNode) (l *BinaryTreeNode, err error) {

	if n.lchild == nil {

		err = fmt.Errorf("the node havn't lchild")
		return l, err

	} else {

		l = n.lchild

		return l, err

	}

}

//获取该节点的Rchild node
func GetRchildNode(n *BinaryTreeNode) (r *BinaryTreeNode, err error) {

	if n.rchild == nil {

		err = fmt.Errorf("the node havn't rchild")
		return r, err

	} else {
		r = n.rchild
		return r, err

	}

}

//获取该节点的parent node
func GetParentNode(n *BinaryTreeNode) (p *BinaryTreeNode, err error) {

	if n.parent == nil {

		err = fmt.Errorf("the root havn't parentnode")
		return p, err

	} else {

		p = n.parent
		return p, err

	}

}

//判断该节点的lchlid node 是否为nil
func IsNullLchildNode(n *BinaryTreeNode) bool {

	if n.lchild == nil {

		return true

	} else {

		return false

	}

}

//判断该节点的rchlid node 是否为nil
func IsNullRchildNode(n *BinaryTreeNode) bool {
	if n.rchild == nil {

		return true

	} else {

		return false

	}
}

//判断该节点的parent node 是否为nil
func IsNullParentNode(n *BinaryTreeNode) bool {
	if n.parent == nil {

		return true

	} else {

		return false

	}
}

//insert lchild node while
func (t *BinaryTree) InsetLchild(p *BinaryTreeNode, l *BinaryTreeNode) bool {

	if p == nil && l == nil {

		return false

	} else if IsNullLchildNode(p) {

		p.lchild = l
		if p.rchild == nil {

			t.height++
			t.size++
			l.parent = p
			return true

		} else {
			t.size++
			l.parent = p
			return true
		}

	} else {

		return false

	}

}

//insert rchild node while
func (t *BinaryTree) InsetRchild(p *BinaryTreeNode, r *BinaryTreeNode) bool {

	if p == nil && r == nil {

		return false

	} else if IsNullRchildNode(p) {

		p.rchild = r
		if p.lchild == nil {

			t.height++
			t.size++
			r.parent = p
			return true

		} else {
			t.size++
			r.parent = p
			return true
		}

	} else {

		return false

	}

}

//set lchild node
func (n *BinaryTreeNode) SetLchild(l *BinaryTreeNode) bool {

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

//rewriting String()string
func (t *BinaryTree) String() string {

	if t.height == 0 || t.size == 0 {

		return "the tree is nil"

	} else {

		l := t.InOrder()
		s := ""
		s = fmt.Sprintf("length:%d.\r\n", l.Len())

		st := "the BinaryTree-"
		st += s
		h := l.Front()
		h1, ok := h.Value.(*BinaryTreeNode)
		if !ok {

			fmt.Println("the typeof not")
			os.Exit(1)

		}
		//reflect.TypeOf(h1)
		st += h1.date
		st += "\r\n"
		//fmt.Println(h1.date)
		s = ""
		for j := 0; j < l.Len()-1; j++ {
			h = h.Next()
			h1 = h.Value.(*BinaryTreeNode)
			s += h1.date
			s += "\r\n"

		}
		st += s

		return st
	}

}

//二叉树的前序排列 中左右
func (t *BinaryTree) PrevOrder() *list.List {

	traversal := list.New()

	prevOrder(t.root, traversal)

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
func (t *BinaryTree) InOrder() *list.List {

	traversal := list.New()

	inOrder(t.root, traversal)

	return traversal

}

//二叉树的中序排列
func inOrder(n *BinaryTreeNode, l *list.List) {

	if n == nil {

		return

	}

	prevOrder(n.lchild, l)
	l.PushBack(n)
	prevOrder(n.rchild, l)

}

//二叉树的后序排列，左右中
func (t *BinaryTree) TailOrder() *list.List {

	traversal := list.New()
	tailOrder(t.root, traversal)
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

type Lister struct {
	L *list.List
}

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

func main() {

	t := MakeTree()
	n2 := &BinaryTreeNode{date: "this is node2"}
	t.root.SetLchild(n2)
	n3 := &BinaryTreeNode{date: "this is node3"}
	t.root.SetRchild(n3)
	n4 := &BinaryTreeNode{date: "this is node4"}
	t.root.lchild.SetLchild(n4)
	n5 := &BinaryTreeNode{date: "this is node5"}
	t.root.lchild.SetRchild(n5)
	n6 := &BinaryTreeNode{date: "this is node6"}
	t.root.rchild.SetLchild(n6)
	n7 := &BinaryTreeNode{date: "this is node7"}
	t.root.rchild.SetRchild(n7)

	l := t.InOrder()
	l1 := &Lister{L: l}
	fmt.Println(l1)

}
