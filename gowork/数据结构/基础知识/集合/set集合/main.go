//set 集合，集合内的元素是无序的且不重复存在的
package main

import (
	"fmt"
)

//set
type HashSet struct {
	m map[interface{}]bool
}

//init
//func (h *HashSet) Init() *HashSet {

//	h.m = make(map[interface{}]bool)
//	return h

//}

func MakeSet() *HashSet {

	h := &HashSet{m: make(map[interface{}]bool)}

	return h

}

//insert value
func (h *HashSet) Insert(it interface{}) bool {

	if _, ok := h.m[it]; ok {

		return false

	}

	h.m[it] = true

	return true

}

//delete value
func (h *HashSet) DelValue(it interface{}) bool {

	if _, ok := h.m[it]; !ok {

		return true

	}

	delete(h.m, it)

	return true

}

//get length
func (h *HashSet) Len() int {

	return len(h.m)

}

//clear hashset
func (h *HashSet) Clear() bool {

	h.m = make(map[interface{}]bool)
	return true

}

//contrain e
func (h *HashSet) Contain(e interface{}) bool {

	return h.m[e]

}

//traversal hashset
func (h *HashSet) Traversal() []interface{} {

	arr := make([]interface{}, h.Len())
	i := 0
	for k, _ := range h.m {

		if i < h.Len() {

			arr[i] = k
			i++
			continue
		}

		//arr = append(arr, k)

	}

	return arr

}

//rewriting String()string
func (h *HashSet) String() string {

	st := fmt.Sprintf("the hashset-length:%d\r\n", h.Len())

	for _, v := range h.Traversal() {

		st += fmt.Sprintf("%v\r\n", v)

	}

	return st

}

func main() {

	h := MakeSet()
	h.Insert("java")
	h.Insert("c++")
	h.Insert("C#")
	h.Insert("golang")
	h.Insert("python")
	h.Insert("php")
	fmt.Println(h)

}
