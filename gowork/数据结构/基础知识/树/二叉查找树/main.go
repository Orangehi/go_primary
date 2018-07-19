//二叉查找树，又称为二叉排序树，二叉搜索树
package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

//若左子树不为空，则左子树上所有的节点均小于或等于他的根节点
//若右子树不为空，则右子树上所有的节点均大于或等于他的根节点
//他的左子树，右子树也分别为二叉查找树
//这种特殊的二叉树，按照中序遍历即可得到一个有序数组
//删除、插入、查找都是O(n)

type (
	Tree struct {
		lchild *Tree
		rchild *Tree
		date   int
	}

	Lister struct {
		li *list.List
	}
)

//构造二叉排序树
func MakeTree(n int, k int) *Tree {

	var t *Tree
	for _, v := range rand.Perm(n) {

		t = InsertNode(t, (1+v)*k)

	}

	return t
}

//插入数据
func InsertNode(t *Tree, v int) *Tree {

	if t == nil {

		return &Tree{nil, nil, v}

	}

	if v < t.date {

		t.lchild = InsertNode(t.lchild, v)
		return t
	}

	t.rchild = InsertNode(t.rchild, v)

	return t

}

//查看树是否包含某值
func Contains(t *Tree, v int) bool {

	if t == nil {

		return false

	}

	switch {
	case t.date > v:
		return Contains(t.lchild, v)
	case t.date < v:
		return Contains(t.rchild, v)
	case t.date == v:
		return true
	}

	return false

}

//搜索数据
func SearchDate(t *Tree, v int) int {

	if t == nil {
		return -2
	}

	switch {
	case t.date == v:
		return t.date
	case t.date > v:
		return SearchDate(t.lchild, v)
	case t.date < v:
		return SearchDate(t.rchild, v)

	}

	return -1
}

//获取最小值
func GetMin(t *Tree) (i int, b bool) {

	if t == nil {

		return -1, false

	}

	for t.lchild != nil {

		t = t.lchild

	}

	return t.date, true

}

//获取最大值
func GetMax(t *Tree) (i int, b bool) {

	if t == nil {

		return -1, false

	}

	for t.rchild != nil {

		t = t.rchild

	}

	return t.date, true

}

//获取节点
func GetNode(t *Tree, v int) *Tree {

	if t == nil {

		return nil

	}

	if !Contains(t, v) {

		return nil

	}

	switch {

	case t.date == v:
		return t
	case t.date > v:
		return GetNode(t.lchild, v)
	case t.date < v:
		return GetNode(t.rchild, v)

	}

	return nil

}

//获取该节点的双亲节点
func GetParentNode(t *Tree, v int) (p *Tree, i int) {

	if !Contains(t, v) {

		return p, -1

	}

	p = t

	if t.lchild != nil && t.lchild.date == v {
		return p, 1
	} else if t.rchild != nil && t.rchild.date == v {
		return p, 2
	}

	if t.date > v {

		return GetParentNode(t.lchild, v)

	}

	if t.date < v {

		return GetParentNode(t.rchild, v)

	}

	return p, -1
}

//删除节点
func DeleteNode(t *Tree, v int) bool {

	n := GetNode(t, v)
	if n == nil {

		return false

	}

	//delete节点分为四种情况

	//该节点为叶子节点，左子树，右子树都为nil
	if n.lchild == nil && n.rchild == nil {

		p, i := GetParentNode(t, n.date)
		if p == nil || i == -1 {

			return false

		} else {

			if i == 1 {

				p.lchild = nil
				n = nil

				return true

			} else {

				p.rchild = nil
				n = nil
				return true
			}

		}

	} else if n.lchild == nil && n.rchild != nil { //该节点的左子树为nil

		p, i := GetParentNode(t, n.date)
		if p == nil || i == -1 {

			return false

		}

		if p.lchild == n {

			p.lchild = n.rchild
			n = nil
			return true

		} else {

			p.rchild = n.rchild
			n = nil
			return true

		}

	} else if n.rchild == nil && n.lchild != nil { //该节点的右子树为nil

		p, i := GetParentNode(t, n.date)
		if p == nil || i == -1 {

			return false

		}

		if p.lchild == n {

			p.lchild = n.lchild
			n = nil
			return true

		} else {

			p.rchild = n.lchild
			n = nil
			return true

		}

	} else { //上述三种情况都不是情况下

		de := GetNode(t, v)
		return Remove(t, de)

	}

}

//当该节点的左右孩子都不为空的时候删除该节点
//只需要将该节点的右子树的最小值和该节点替换，然后删除该节点的右子树的最小值，即可完成删除
func Remove(t *Tree, n *Tree) bool {
	//获取该节点的双亲节点
	p, _ := GetParentNode(t, n.date)

	//获取该节点右子树的最小值的date
	n_Min_v, _ := GetMin(n.rchild)
	//获取该节点的右子树的最小值
	n_Min := GetNode(t, n_Min_v)

	//判断删除的该节点是其双亲节点的左孩子还是右孩子
	if n.date > p.date { //该节点是其双亲节点的右孩子

		if n.rchild == n_Min { //如果该节点的右孩子为最小值

			p.rchild = n_Min
			n_Min.lchild = n.lchild
			n = nil
			return true

		} else {

			p.rchild = n_Min
			if !DeleteNode(n.rchild, n_Min_v) {
				return false

			}
			p.rchild.lchild = n.lchild
			p.rchild.rchild = n.rchild

			n = nil
			return true

		}

	} else { //该节点是其双亲节点的左孩子

		if n.rchild == n_Min { //如果该节点的右孩子为最小值

			p.lchild = n_Min

			n_Min.lchild = n.lchild

			n = nil
			return true

		} else {

			p.lchild = n_Min
			DeleteNode(n.rchild, n_Min_v)
			p.lchild.lchild = n.lchild
			p.lchild.rchild = n.rchild

			n = nil

			return true

		}

	}

}

//中序化
func (t *Tree) InOrder() *list.List {

	l := list.New()

	inOrder(t, l)

	return l

}

//中序遍历
func inOrder(t *Tree, li *list.List) {

	if t == nil {

		return

	}
	inOrder(t.lchild, li)
	li.PushBack(t)
	inOrder(t.rchild, li)

}

//重写双链表List打印方法
func (lis Lister) String() string {

	h := lis.li.Front()
	h1 := h.Value.(*Tree)

	st := fmt.Sprintf("the tree- length:%d\r\n", lis.li.Len())
	//st += fmt.Sprintf("the date:%d\r\n", h1.date)
	for i := 0; i < lis.li.Len(); i++ {
		h1 = h.Value.(*Tree)
		st += fmt.Sprintf("the date:%d\r\n", h1.date)
		h = h.Next()

	}

	return st

}

//前序化
func (t *Tree) PreOrder() *list.List {

	l := list.New()

	preOrder(t, l)

	return l

}

//前序遍历
func preOrder(t *Tree, li *list.List) {

	if t == nil {

		return

	}

	li.PushBack(t)
	preOrder(t.lchild, li)

	preOrder(t.rchild, li)

}

//后序化
func (t *Tree) TailOrder() *list.List {

	l := list.New()

	tailOrder(t, l)

	return l

}

//后序遍历
func tailOrder(t *Tree, li *list.List) {

	if t == nil {

		return

	}

	tailOrder(t.lchild, li)

	tailOrder(t.rchild, li)
	li.PushBack(t)
}

//中序遍历
func InPrintf(t *Tree) {

	if t == nil {

		return
	}

	InPrintf(t.lchild)
	fmt.Print(t.date)
	InPrintf(t.rchild)

}

//前序遍历
func PrevPrintf(t *Tree) {

	if t == nil {

		return
	}
	fmt.Println(t.date)
	PrevPrintf(t.lchild)
	PrevPrintf(t.rchild)

}

//后序遍历
func TailPrintf(t *Tree) {

	if t == nil {

		return
	}

	TailPrintf(t.lchild)
	TailPrintf(t.rchild)
	fmt.Print(t.date)
}

func main() {

	t := MakeTree(20, 1)

	InPrintf(t)
	fmt.Println()
	DeleteNode(t, 16)
	InPrintf(t)

	//	n := GetNode(t, 10)
	//	if n.lchild != nil {

	//		fmt.Println("n.lchild", n.lchild.date)

	//	}
	//	if n.rchild != nil {

	//		fmt.Println("n.rchild", n.rchild.date)

	//	}

	//	fmt.Println(GetMin(n))

	//	fmt.Println(SearchDate(t, 11))
	//	fmt.Println(GetMin(t))
	//	fmt.Println(GetMax(t))
	//	fmt.Println(Contains(t, 3))

	//	lis := t.TailOrder()
	//	l := Lister{}
	//	l.li = lis
	//	fmt.Println(l)

}
