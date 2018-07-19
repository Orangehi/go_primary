//线索二叉树（只有lchild指针,以及rchild指针，不带parent指针）
//假设二叉树存在n个节点，那么将存在n+1个空指针,n-1条分支线
package main

import (
	"container/list"
	"fmt"
)

//线索二叉树的node
type BinaryTreeNode struct {
	date string
	//	parent         *BinaryTreeNode //parent node
	lchild, rchild *BinaryTreeNode
	lch, rch       bool //lch == true 说明lchild为线索指针，否则lchild执行该节点的左孩子
	//rch == true 说明rchild为线索指针，否则rchild执行该节点的右孩子
}

var preNode *BinaryTreeNode

//construt function
func MakeTree() *BinaryTreeNode {

	n := &BinaryTreeNode{date: "this is root", lch: false, rch: false}
	return n

}

//set lchild node
func (n *BinaryTreeNode) SetLchild(l *BinaryTreeNode) bool {

	if n.lchild != nil {

		return false
	}

	n.lchild = l

	//l.parent = n
	return true

}

//set rchild node
func (n *BinaryTreeNode) SetRchild(r *BinaryTreeNode) bool {

	if n.rchild != nil {

		return false
	}

	n.rchild = r
	//r.parent = n
	return true

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

//二叉树中序线索化二叉树
func InOrderTree(n *BinaryTreeNode) {

	if n == nil {

		return
	}

	InOrderTree(n.lchild) //中序线索化左子树

	//左叶子节点中序线索化
	if n.lchild == nil {

		n.lchild = preNode
		n.lch = true

	}

	//右叶子节点中序线索化
	if preNode != nil && preNode.rchild == nil {

		preNode.rchild = n
		preNode.rch = true

	}

	preNode = n //保存当前节点

	InOrderTree(n.rchild) //中序线索化右子树

}

//中序遍历线索二叉树,按照后继节点遍历,正序排列
func TraversalInTreePre(t *BinaryTreeNode) {

	InOrderTree(t)

	n := t

	for n.lchild != nil && !n.lch {

		n = n.lchild

	}

	for n != nil {
		//如果当前节点不为空，输出当前节点
		fmt.Println(n.date)

		//如果当前节点的右孩子为后继线索指针，输出当前节点的右孩子
		if n.rch {

			n = n.rchild

			//如果当前节点的右孩子不为后继线索指针，找到当前节点的最左孩子
		} else {

			n = n.rchild

			for n != nil && !n.lch {

				n = n.lchild

			}

		}

	}

}

//中序遍历线索二叉树,按照前驱节点遍历，倒序排列
func TraversalInTreeTail(t *BinaryTreeNode) {

	InOrderTree(t)

	n := t

	for n.rchild != nil && !n.rch {

		n = n.rchild

	}

	for n != nil {

		//如果当前节点不为空，输出当前节点
		fmt.Println(n.date)

		//如果当前节点的左孩子为线索指针，输出当前节点的左孩子
		if n.lch {

			n = n.lchild

			//如果当前节点的左孩子不为线索指针，找到当前节点的最右孩子
		} else {

			n = n.lchild

			for n != nil && !n.rch {

				n = n.rchild

			}

		}

	}

}

//前序化线索二叉树
func PrevOrderTree(n *BinaryTreeNode) {

	if n == nil {
		return
	}

	if n.lchild == nil {

		n = preNode
		n.lch = true

	}

	if preNode != nil && n.rchild == nil {

	}

	PrevOrderTree(n.lchild)

	PrevOrderTree(n.rchild)

}

func main() {

	t := MakeTree()

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

	//TraversalInTreePre(t)
	TraversalInTreeTail(t)

	//	l := t.InOrder()
	//	l1 := &Lister{L: l}
	//	fmt.Println(l1)
}
