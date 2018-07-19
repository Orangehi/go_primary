//双循环链表：双链表的基础再加上尾节点指向头结点这一特点,头结点指向尾节点
package main

import (
	"fmt"
)

type Mylist struct {
	date   string
	top    *Dnode
	length int
}

type Dnode struct {
	value string
	last  *Dnode
	next  *Dnode
}

//construct function
func Makelist() *Mylist {

	m := Mylist{date: "this is top"}
	return &m

}

//insert from head
func (m *Mylist) InsertHead(n *Dnode) bool {

	if m.top == nil {

		m.top = n
		m.length++
		return true

	} else if m.length == 1 {

		h := m.top
		m.top = n
		m.top.last = h
		m.top.next = h
		h.last = m.top
		h.next = m.top

		m.length++
		return true

	} else {
		h := m.top
		m.top = n
		n.next = h
		n.last = h.last
		h.last = n

		m.length++
		return true
	}

}

//insert from tail
func (m *Mylist) InsertTail(n *Dnode) bool {

	if m.top == nil {

		m.top = n
		m.length++
		return true

	} else if m.length == 1 {

		m.top.next = n
		m.top.last = n
		n.next = m.top
		n.last = m.top
		m.length++
		return true

	} else {

		h := m.top
		for j := 0; j < m.length-1; j++ {

			h = h.next

		}
		h.next = n
		n.last = h
		n.next = m.top
		m.top.last = n

		m.length++
		return true
	}

}

//insert from int
func (m *Mylist) Insertint(i int, n *Dnode) bool {
	if i <= 0 || i > m.length+1 {

		return false

	} else if i == 1 {
		m.InsertHead(n)
		return true
	} else if i == m.length+1 {
		m.InsertTail(n)
		return true

	} else if i > 1 || i <= m.length {

		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}
		h1 := h.next
		h.next = n
		n.last = h
		n.next = h1
		h1.last = n

		m.length++
		return true
	} else {
		return false
	}

}

//delete from int
func (m *Mylist) DelDnode(i int) bool {

	if i <= 0 || i > m.length+1 {

		return false

	} else if i > 1 && i == m.length {

		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}
		h1 := h.next
		h.next = nil
		h.next = h1.next
		h1.next.last = h
		m.length--
		return true

	} else if i == 1 && m.length == 1 {

		m.top = nil
		m.length--
		return true

	} else if i == 1 && m.length > 1 {

		m.length--
		return true

	} else if i > 1 && i < m.length {
		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}
		h1 := h.next
		h.next = nil
		h.next = h1.next
		h1.next.last = h
		m.length--
		return true

		m.length--
		return true

	} else {
		return false
	}

}

//clear
func (m *Mylist) Clear() bool {

	if m.length == 0 {

		return true

	} else if m.length == 1 {

		m.top = nil
		m.length = 0
		return true

	} else {
		for m.length > 0 {

			m.DelDnode(m.length)

		}
		return true
	}

}

//select from int
func (m *Mylist) Selint(i int) (n *Dnode, err error) {

	if m.length == 0 {
		err = fmt.Errorf("the list is nil")

		return n, err

	}

	if i <= 0 || i > m.length {

		err = fmt.Errorf("please input num>=0 and num < length")

		return n, err

	} else {

		h := m.top
		for j := 0; j < i-1; j++ {

			h = h.next

		}

		return h, err

	}

}

//rewriting String()string
func (m *Mylist) String() string {

	if m.length == 0 {

		return "this is nil \r\nlength:0"

	} else if m.length == 1 {

		st := "the list: "
		s := ""
		s = fmt.Sprintf("id:%d value:%s;", 1, m.top.value)
		st += s
		s = fmt.Sprintf("length:%d", m.length)
		st += "\r\n"
		st += s
		return st

	} else {

		st := "the list: "
		i := 1
		s := ""
		s = fmt.Sprintf("id:%d value:%s;", i, m.top.value)
		st += s
		s = ""
		h := m.top
		for j := 0; j < m.length-1; j++ {
			h = h.next
			i++
			s += fmt.Sprintf("id:%d value:%s;", i, h.value)

		}

		st += s
		s = fmt.Sprintf("length:%d", m.length)
		st += "\r\n"
		st += s
		return st
	}

}

func main() {

	m := Makelist()
	fmt.Println(m)

	n1 := &Dnode{value: "this is node1"}
	m.InsertTail(n1)
	fmt.Println(m)

	n2 := &Dnode{value: "this is node2"}
	m.InsertTail(n2)
	fmt.Println(m)

	n3 := &Dnode{value: "this is node3"}
	m.Insertint(2, n3)
	fmt.Println(m)

	n4 := &Dnode{value: "this is node4"}
	m.Insertint(1, n4)
	fmt.Println(m)

	n5 := &Dnode{value: "this is node5"}
	m.Insertint(5, n5)
	fmt.Println(m)

	n6, err := m.Selint(2)
	if err != nil {

		fmt.Println(err)

	} else {

		fmt.Println("id:", 2, n6.value)

	}

	//	m.DelDnode(2)
	//	fmt.Println()
	//	fmt.Println(m)
	//	m.DelDnode(2)
	//	fmt.Println()
	//	fmt.Println(m)
	//	m.DelDnode(2)
	//	fmt.Println()
	//	fmt.Println(m)
	//	m.DelDnode(2)
	//	fmt.Println()
	//	fmt.Println(m)
	//	m.DelDnode(1)
	//	fmt.Println()
	//	fmt.Println(m)
	m.Clear()

}
