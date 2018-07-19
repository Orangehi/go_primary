//单链表,带头结点
package main

import (
	"fmt"
)

//头结点
type Mylist struct {
	id   int
	head *Node
}

//链表节点
type Node struct {
	value []string
	next  *Node
}

//头结点初始化
func Makelist() *Mylist {

	m := Mylist{id: 0, head: &Node{value: []string{"this", "is", "head"}}}
	return &m

}

//插入一个链表到头部
func (N *Mylist) inserttop(n *Node) bool {

	h := N.head.next

	N.head.next = n

	N.head.next.next = h

	return true

}

//插入一个链表到尾部
func (N *Mylist) insertlow(n *Node) bool {

	h := N.head

	for h.next != nil {

		h = h.next

	}

	h.next = n

	return true

}

//判断单列表是否为空
func (n *Mylist) Isnull() bool {

	return n.head.next == nil

}

//判断链表的长度
func (n *Mylist) Length() (i int) {

	if n.head.next == nil {

		return i

	}
	h := n.head

	for h.next != nil {

		i++

		h = h.next

	}
	return i
}

//清空节点
func (N *Mylist) Clear() bool {

	n := N.head.next
	for n != nil {

		tmp := n

		n = nil

		n = tmp.next

	}

	N.head.next = nil

	return true

}

//插入到指定节点
func (N *Mylist) insertint(i int, n *Node) bool {

	if i < 0 || i > N.Length() {

		return false

	}

	h := N.head

	for j := 0; j < (i - 1); j++ {

		h = h.next

	}

	//	elem.next = current.next //将后边的链，放到的elem current.next = &elem

	n.next = h.next
	h.next = n

	return true

}

//删除节点
func (N *Mylist) delnode(i int) bool {

	if i < 0 || i > N.Length() {
		return false
	}

	h := N.head

	for j := 0; j < (i - 1); j++ {

		h = h.next

	}

	h.next = h.next.next
	return true

}

func listTos(vs []string) string {

	s := ""
	for _, v := range vs {

		s += v

	}

	return s
}

//重写String（）
func (N *Mylist) String() string {

	if N.Length() == 0 {

		return "this is null"

	}

	h := N.head.next

	i := 0

	st := "链表内容:"

	for h != nil {
		i++
		s := fmt.Sprintf("%d:%s;", i, listTos(h.value))
		st += s
		h = h.next

	}

	return st

}

//获取指定位置，节点的值
func (N *Mylist) getvs(i int) string {

	if i < 0 || i > N.Length() {

		return "参数无效"

	}

	h := N.head

	for j := 0; j < i-1; j++ {

		h = h.next

	}

	return listTos(h.next.value)

}

func main() {

	m := Makelist()
	n1 := &Node{value: []string{"this", "is", "node1"}}
	m.insertlow(n1)
	fmt.Println(m.Isnull(), m.head.next.value)
	n2 := &Node{value: []string{"this", "is", "node2"}}
	m.insertlow(n2)
	n3 := &Node{value: []string{"this", "is", "node3"}}
	m.insertlow(n3)
	fmt.Println(m.Isnull(), m.head.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.next.value)
	fmt.Println(m.Length())
	n4 := &Node{value: []string{"this", "is", "node4"}}
	fmt.Println(m.insertint(2, n4))
	fmt.Println(m.Length())
	fmt.Println(m.Isnull(), m.head.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.next.next.value)
	fmt.Println(m.delnode(2))
	fmt.Println(m.Length())
	fmt.Println(m.Isnull(), m.head.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.value)
	fmt.Println(m.Isnull(), m.head.next.next.next.value)
	fmt.Println("获取:", m.getvs(2))
	fmt.Println(m)
	//fmt.Println(m.Isnull(), m.head.next.next.next.next.value)
}
