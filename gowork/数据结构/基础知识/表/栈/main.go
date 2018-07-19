//栈-先进后出
package main

import (
	"fmt"
)

type stack struct {
	size  int
	top   int
	value []string
}

//构造新栈
func Makestack(size int) stack {

	s := stack{}
	s.size = size
	s.value = make([]string, size)

	return s

}

//向栈存放数据
func (t *stack) push(date string) bool {

	if t.isfull() {
		return false
	}

	t.value[t.top] = date

	t.top++

	t.size++

	return true

}

//向栈取出数据
func (t *stack) pop() (date string, err error) {

	if t.isnull() {

		date = "无"
		err = fmt.Errorf("：栈为空")

		return date, err

	}
	t.top--

	date = t.value[t.top]

	err = fmt.Errorf("")

	return

}

//判断栈是否满
func (t *stack) isfull() bool {

	if t.top == t.size {

		return true

	} else {
		return false
	}

}

//判断栈是否为空
func (t *stack) isnull() bool {

	if t.top == 0 {
		return true
	} else {
		return false
	}

}

func main() {

	s1 := Makestack(10)

	s1.push("hello")
	s1.push("world")
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())

}
