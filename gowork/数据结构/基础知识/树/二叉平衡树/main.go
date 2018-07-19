//二叉平衡树是二叉查找树的优化
package main

import (
	"fmt"
	"math/rand"
)

//二叉平衡树的插入、删除、查找都是O(logn)
//其左子树及其右子树都是高度平衡的二叉树
//左子树和右子树的高度之差（深度）的绝对值不超过1.
//尽量使其上的节点距离根节点越近，提示增删查的速率

//二叉平衡树
type AVLNode struct {
	date    int //节点的值
	balance int //平衡因子
	//	parent  *AVLNode //双亲节点
	lchild *AVLNode //左孩子
	rchild *AVLNode //右孩子
}

func NewTree() *AVLNode {

	var t *AVLNode
	//t := new(AVLNode)

	for _, v := range rand.Perm(20) {

		t = InsertNode(t, (v+1)*1)

	}

	return t
}

//获取树的深度 balance
func GetBal(t *AVLNode) int {

	if t == nil {

		return -1

	}

	return t.balance

}

//返回一个数的绝对值
func abs(i int) int {

	if i < 0 {

		return -i

	}
	if i == 0 {

		return 0

	}

	return i

}

//判断该节点是否平衡
func IsBalance(p *AVLNode) bool {

	n := GetBal(p.lchild) - GetBal(p.rchild)

	n1 := abs(n)

	if n1 == 1 || n1 == 0 {

		return true

	} else {

		return false

	}

}

//判断两个值的大小
func Max(i int, n int) int {

	if i > n {

		return i

	} else {

		return n

	}

}

//二叉平衡树的插入
//插入主要分为四种
func InsertNode(root *AVLNode, value int) *AVLNode {

	if root == nil { //递归的退出条件，如果该节点为空，返回一个date= value的节点

		root = &AVLNode{date: value}
		return root

	}

	if value > root.date { //插入到该节点的右边

		root.rchild = InsertNode(root.rchild, value) //向该节点的右孩子插入date= value的节点

		if !IsBalance(root) { //判断该节点是否平衡,如果不平衡执行下列操作

			if value > root.rchild.date { //判断插入的是否为该节点的右孩子的右子树,如果是，对该节点进行rr平衡操作

				root = right_right_Insert(root) //rr平衡操作

			} else { //判断插入的是否是该节点的右孩子的左子树，如果是，对该节点进行rl平衡操作

				root = right_left_Insert(root) //rl平衡操作

			}

		}

	} else { //插入到该节点的左边

		root.lchild = InsertNode(root.lchild, value) //向该节点的左孩子插入date = value的节点
		if !IsBalance(root) {                        //判断该节点是否平衡，如果不平衡执行下列操作

			if value > root.lchild.date { //判断插入的节点是否是该节点的左孩子的右子树，如果是，对该节点进行lr操作

				root = left_right_Insert(root)

			} else { //判断插入的节点是否是该节点的左孩子的左子树，如果是 ，对该节点进行ll操作

				root = left_left_Insert(root)

			}

		}

	}

	root.balance = Max(GetBal(root.lchild), GetBal(root.rchild)) + 1

	return root

}

//对该节点的左孩子的左子树进行插入
func left_left_Insert(p *AVLNode) *AVLNode {

	t := p.lchild
	p.lchild = t.rchild
	t.rchild = p

	p.balance = Max(GetBal(p.lchild), GetBal(p.rchild)) + 1
	t.balance = Max(GetBal(t.lchild), GetBal(t.rchild)) + 1

	return t

}

//对该节点的左孩子的右子树进行插入
func left_right_Insert(p *AVLNode) *AVLNode {

	t := p.lchild
	p.lchild = right_right_Insert(t)

	return left_left_Insert(p)

}

//对该节点的右孩子的左子树进行插入
func right_left_Insert(p *AVLNode) *AVLNode {

	t := p.rchild
	p.rchild = left_left_Insert(t)

	return right_right_Insert(p)

}

//对该节点的右孩子的右子树进行插入
func right_right_Insert(p *AVLNode) *AVLNode {

	t := p.rchild
	p.rchild = t.lchild
	t.lchild = p

	p.balance = Max(GetBal(p.lchild), GetBal(p.rchild)) + 1
	t.balance = Max(GetBal(t.lchild), GetBal(t.rchild)) + 1

	return t

}

//中序遍历二叉平衡树
func InOrder(root *AVLNode) {

	if root == nil {

		return

	}

	InOrder(root.lchild)
	fmt.Println(root.date)
	InOrder(root.rchild)

}

//删除二叉平衡书的节点
func DelNode(root *AVLNode, value int) {

}

func main() {

	t := NewTree()
	InOrder(t)

	//fmt.Println(len(rand.Perm(19)))

}
