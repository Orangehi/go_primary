//顺序表 插入或者删除一个（除去尾部的）数据，会影响整个表的排序
package main

import (
	"fmt"
)

const Max = 20

type arr struct {
	date   [Max]string
	length int
}

func Makearr() *arr {

	a := arr{}

	return &a

}

//向顺序表中插入一个数据
func (a *arr) insert(i int, value string) bool {

	if i <= 0 || i > Max || a.length > Max {

		return false

	}

	if i < a.length {

		st := a.date

		for j := i - 1; j < a.length; j++ {

			if st[j] == "" {
				continue
			}

			a.date[j+1] = st[j]

		}
		a.date[i-1] = value
		a.length++
		return true

	} else {

		a.date[i-1] = value
		a.length++
		return true

	}

}

//重写String()
func (a *arr) String() string {
	st := "arr: "
	s := ""
	for i, v := range a.date {

		if v == "" {

			continue

		}

		s += fmt.Sprintf("%d:%s;", i+1, v)

	}

	st += s

	return st

}

//删除顺序表中元素
func (a *arr) Dele(i int) bool {

	if i <= 0 || i > Max || i > a.length {

		return false

	}

	if i == a.length {

		a.date[i-1] = ""
		a.length--

		return true

	} else {

		a.date[i-1] = ""
		st := a.date
		a.length--
		return true

		for j := i - 1; j < a.length-1; j++ {

			a.date[j] = st[j+1]

		}
		a.date[i-1] = ""
		a.length--
		return true

	}

}

//获取顺序表中的元素
func (a *arr) get(i int) string {
	if i <= 0 || i > Max || i > a.length {

		return false

	}

	return a.date[i]

}

func main() {

	a := Makearr()

	a.insert(1, "java")
	fmt.Println(a.length)
	a.insert(2, "c++")
	fmt.Println(a.length)
	a.insert(3, "c")
	//	fmt.Println(a.length)
	a.insert(4, "c#")
	//	fmt.Println(a.length)
	//	a.insert(5, "go")
	//	a.insert(6, "rul")
	a.insert(1, "php")
	//	a.insert(1,"java")
	//	a.insert(1,"java")

	fmt.Println(a)

	a.Dele(5)
	fmt.Println(a)

	a.Dele(2)
	fmt.Println(a)
}
