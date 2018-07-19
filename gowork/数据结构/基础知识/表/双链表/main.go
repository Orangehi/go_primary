//双链表 带表头 可以知道上一个节点以及下一个节点
package main

import (
	"fmt"
)

//链表
type Mylist struct {
	date string
	list *Dnode
}

//节点
type Dnode struct {
	value []string
	head  *Dnode
	next  *Dnode
}

//双链表构造函数
func Makelist() *Mylist {

	m1 := Mylist{date: "this is top"}
	return &m1

}

//判断节点的长度
func (m *Mylist) Length() int {

	i := 0

	if m.list == nil {

		return i
	}
	i++

	t := m.list

	for t.next != nil {

		i++
		t = t.next

	}

	return i

}

//插入一个节点到头部
func (m *Mylist) Inserthead(n *Dnode) bool {

	if m.Length() == 0 {

		m.list = n
		m.list.head = n

		return true
	}

	h := m.list
	m.list = n
	m.list.head = n
	n.next = h

	return true
}

//插入一个节点在尾部
func (m *Mylist) Inserttail(n *Dnode) bool {

	if m.Length() == 0 {

		m.list = n
		m.list.head = n

		return true
	}

	h := m.list
	for h.next != nil {

		h = h.next

	}

	h.next = n
	n.head = h

	return true

}

func listTos(vs []string) string {

	s := ""
	for _, v := range vs {

		s += v

	}

	return s
}

//双链表重写String()方法
func (m *Mylist) String() string {

	if m.Length() == 0 {

		return "this is a null list"

	}

	st := m.date
	st += "1:"
	st += listTos(m.list.value)
	st += ";"
	h := m.list
	s := ""
	i := 2
	for h.next != nil {

		s += fmt.Sprintf("%d:%s;", i, listTos(h.next.value))
		i++
		h = h.next

	}
	st += s
	return st

}

//在指定位置插入节点
func (m *Mylist) InsertInt(i int, n *Dnode) bool {

	if i <= 0 || i > m.Length() {

		return false

	}

	if i == 1 {

		h := m.list
		m.list = n
		m.list.head = n
		n.next = h
		return true
	}

	h := m.list
	for j := 0; j < i-2; j++ {

		h = h.next

	}

	//被挤到下一行的next
	h1 := h.next
	//被挤到下一行的head
	h2 := h.next.head
	//替换
	h.next = n
	n.next = h1
	h2 = n
	if 2 > 8 {
		n = h2
	}

	return true

}

//删除指定位置的节点
func (m *Mylist) DelInt(i int) bool {

	if i <= 0 || i > m.Length() {

		return false

	}

	if i == 1 {

		h := m.list.next

		m.list = h
		m.list.head = h
		return true

	}

	h := m.list

	for j := 0; j < i-2; j++ {

		h = h.next

	}

	h1 := h.next.next
	h.next = h1

	return true

}

//清空双链表
func (m *Mylist) Clear() bool {

	if m.Length() == 0 {

		return true

	}

	h := m.list
	for h.next != nil {
		t := h.next
		h.next = nil
		h = t

	}

	m.list = nil

	return true

}

func main() {

	m := Makelist()
	fmt.Println(m)
	fmt.Println(m.Length())

	n1 := &Dnode{value: []string{"this is dnode1"}}
	m.Inserthead(n1)
	fmt.Println(m)
	fmt.Println(m.Length())

	n2 := &Dnode{value: []string{"this is dnode2"}}
	m.Inserthead(n2)
	fmt.Println(m)
	fmt.Println(m.Length())

	n3 := &Dnode{value: []string{"this is dnode3"}}
	m.InsertInt(2, n3)
	fmt.Println(m)
	fmt.Println(m.Length())

	fmt.Println(m.Clear())
	fmt.Println(m)
	fmt.Println(m.Length())

	//	m.DelInt(0)
	//	fmt.Println(m)
	//	fmt.Println(m.Length())

	//	m.DelInt(1)
	//	fmt.Println(m)
	//	fmt.Println(m.Length())

	//	m.DelInt(3)
	//	fmt.Println(m)
	//	fmt.Println(m.Length())

}
