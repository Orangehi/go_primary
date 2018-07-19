//环形队列，先进先出，追加至对尾，弹出对顶
package main

import (
	"fmt"
)

//环形队列
type queen struct {
	length int
	cap    int
	tail   int
	head   int
	value  []string
}

//环形队列构造函数
func Makequeen(size int) queen {

	q := queen{}
	q.length = size
	q.value = make([]string, size)

	return q

}

//环形队列判断为空
func (q *queen) Isnull() bool {

	return q.cap == 0

}

//环形队列判断已满
func (q *queen) Isfull() bool {

	return q.cap == q.length

}

//环形队列添加元素
func (q *queen) push(date string) bool {

	if q.Isfull() {

		fmt.Println("error")
		return false
	}

	q.value[q.tail] = date
	q.tail++
	q.cap++

	return true
}

//环形队列取出元素
func (q *queen) pop() (date string, err error) {

	if q.Isnull() {

		date = "无"

		err = fmt.Errorf(":没有元素")

		return

	}

	date = q.value[q.head]
	q.cap--
	q.head++

	err = fmt.Errorf("")
	return
}

func main() {

	q := Makequeen(10)

	q.push("1")
	q.push("2")
	q.push("3")
	q.push("4")
	q.push("5")
	q.push("6")
	q.push("7")
	q.push("8")
	q.push("9")
	q.push("10")
	q.push("11")

	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

}
