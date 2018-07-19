//单循环链表，可以从任意节点位置访问其它节点，但是只能从当前节点访问后继节点
package main

import (
	"fmt"
)

//top
type Mylist struct {
	date string
	top  *Dnode
}

type Dnode struct {
	value []string
	next  *Dnode
}

//construct function
func Makelist() *Mylist {

	m := Mylist{date: "this is top"}
	return &m

}

//insert into node from top
func (m *Mylist) Insertop(n *Dnode) bool {

	if m.Length() == 0 {

		m.top = n
		m.top.next = m.top
		return true

	}

	h := m.top
	n.next = h
	for h.next != m.top {

		h = h.next

	}
	h.next = n
	m.top = n

	return true

}

//Mylist's length
func (m *Mylist) Length() int {

	i := 0
	if m.top == nil {

		return i

	} else if m.top == m.top.next {

		i = 1

		return i

	} else {
		i = 2
		h := m.top.next
		for h.next != m.top {
			i += 1
			h = h.next
		}
		return i

	}

}

//[]string to string
func Stos(s []string) string {

	st := ""
	for _, v := range s {

		st += v

	}

	return st

}

//prinit rewriting String()
func (m *Mylist) String() string {

	if m.Length() == 0 {

		return "this is nil"

	}

	if m.Length() == 1 {

		//h := m.top
		st := m.date
		st += ": 1:"
		st += Stos(m.top.value) + ";"
		s := fmt.Sprintf(" length:%d", m.Length())
		st += s
		return st

	}

	h := m.top.next
	st := m.date
	st += ": id:1 :"
	st += Stos(m.top.value) + ";"
	s := ""
	i := 1
	for h.next != m.top {
		i++
		s += fmt.Sprintf("id:%d :%s;", i, Stos(h.value))
		h = h.next

	}
	i++
	s += fmt.Sprintf("id:%d :%s;", i, Stos(h.value))
	st += s
	s = fmt.Sprintf(" length:%d", m.Length())
	st += s
	return st

}

//insert node from tail
func (m *Mylist) Insertail(n *Dnode) bool {

	if m.Length() == 0 {

		m.top = n
		n.next = m.top
		return true

	}

	h := m.top
	for h.next != m.top {

		h = h.next

	}
	h.next = n
	n.next = m.top

	return true

}

//insert node from int
func (m *Mylist) Insertint(i int, n *Dnode) bool {

	if i <= 0 || i > (m.Length()+1) {

		return false

	} else if i == 1 {

		return m.Insertop(n)

	} else if i == m.Length()+1 {

		return m.Insertail(n)

	} else {
		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}

		h1 := h.next
		h.next = n
		n.next = h1

		return true

	}

}

//delete node from int
func (m *Mylist) DelNode(i int) bool {

	if i <= 0 || i > m.Length() {

		return false
	} else if i == m.Length() {

		if m.Length() == 1 {

			m.top = nil
			return true

		} else {

			h := m.top
			for j := 0; j < i-2; j++ {

				h = h.next

			}
			h.next = nil
			h.next = m.top

			return true

		}

	} else if i == 1 {

		h := m.top
		h1 := m.top
		for h.next != m.top {

			h = h.next

		}

		m.top = nil
		m.top = h1.next
		h.next = m.top
		return true

	} else {

		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}
		h1 := h.next
		h.next = nil
		h.next = h1.next

		return true

	}

}

//clear Mylist
func (m *Mylist) Clear() bool {

	if m.Length() == 0 {

		return true

	} else if m.Length() == 1 {

		m.top = nil

		return true
	} else {
		h := m.top
		//h1:=m.top
		for h.next != m.top {

			h1 := h
			h = nil
			h = h1.next

		}
		m.top = nil

		return true
	}

}

//select node
func (m *Mylist) SelInt(i int) (n *Dnode, err error) {

	if i <= 0 || i > m.Length() {

		err = fmt.Errorf("this is nil")

		return nil, err

	} else if i == 1 {

		return m.top, err

	} else {

		h := m.top
		for j := 0; j < i-2; j++ {

			h = h.next

		}
		n = h.next

		return n, err

	}

}

func main() {

	m := Makelist()
	fmt.Println(m)

	n1 := &Dnode{value: []string{"this is node1"}}
	m.Insertint(1, n1)
	fmt.Println(m)

	n2 := &Dnode{value: []string{"this is node2"}}
	m.Insertint(2, n2)
	fmt.Println(m)

	n3 := &Dnode{value: []string{"this is node3"}}
	m.Insertint(2, n3)
	fmt.Println()
	fmt.Println(m)

	//	m.Clear()
	//	fmt.Println(m.Length())
	//	fmt.Println(m)

	n4 := &Dnode{value: []string{"this is node4"}}
	m.Insertint(2, n4)
	fmt.Println()
	fmt.Println(m)

	//	n, err := m.SelInt(4)
	//	if err != nil {

	//		fmt.Println(err)

	//	} else {
	//		fmt.Println(n.value)
	//	}

	h := m.top
	h = h.top.next
	fmt.Println(h)

	//m.Clear()
	fmt.Println(m.Length())
	fmt.Println(m)

	//	n5 := &Dnode{value: []string{"this is node5"}}
	//	m.Insertail(n5)
	//	fmt.Println()
	//	fmt.Println(m)

}
