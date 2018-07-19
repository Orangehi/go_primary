//利用go语言字典map来实现set集合。（set集合：集合内元素排序不规律，集合内元素不能出现相同。）
package set

import (
	"bytes"
	"fmt"
)

//Hashset 结构体
type Hashset struct {
	m map[interface{}]bool
}

//Hashset 初始化
func Newhashset() *Hashset {

	return &Hashset{m: make(map[interface{}]bool)}

}

//向Hashset添加元素
func (set *Hashset) Add(e interface{}) bool {

	if !set.m[e] {

		set.m[e] = true

		return true
	}

	return false

}

//删除元素
func (set *Hashset) Delete(e interface{}) {

	delete(set.m, e)

}

//清空元素
func (set *Hashset) Clear() {

	set.m = make(map[interface{}]bool)

}

//元素确认
func (set *Hashset) Contant(e interface{}) bool {

	return set.m[e]

}

//查看hashset集合的长度
func (set *Hashset) Len() int {

	return len(set.m)

}

//Hashset与其他Hashset比较
func (set *Hashset) Same(other *Hashset) bool {

	if other == nil {

		return false

	}

	if set.Len() != other.Len() {

		return false

	}

	for key := range set.m {

		if !other.Contant(key) {

			return false

		}

	}

	return true

}

//遍历Hashset集合
func (set *Hashset) Elements() []interface{} {

	ini := len(set.m)

	snap := make([]interface{}, ini)

	act := 0

	for key := range set.m {

		if act < ini {

			snap[act] = key

		} else {

			snap = append(snap, act)

		}

		act++

	}

	if act < ini {

		snap = snap[:act]

	}

	return snap

}

func (set *Hashset) String() string {

	var buf bytes.Buffer

	buf.WriteString("Set{")

	first := true

	for key := range set.m {

		if first {

			first = false

		} else {

			buf.WriteString(" ")

		}

		buf.WriteString(fmt.Sprintf("%v", key))

	}

	buf.WriteString("}")

	return buf.String()
}
