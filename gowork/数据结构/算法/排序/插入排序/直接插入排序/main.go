// insertion sort 插入排序，时间复杂度为O(n^2)
package main

import (
	"fmt"
)

//插入排序原理：将数组中的元素依次有序的插入到一个有序数组中，从而得到一个有序数组
//插入排序步骤：
//首先取得数组第一个数据，这样就得到一个有序数组。
//然后将数组第二个数据有序的插入到这个数组中，这样就得到一个有序数组。
//依次插入原数组中剩余的数据。
//当将源数组最后一个数据有序插入到有序数组中时，即完成了该数组的排序。

func main() {

	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8}
	InsertSort(s)
	fmt.Println(s)

}

//插入排序
func InsertSort(s []int) {
	j := 0
	t := 0
	for i := 1; i < len(s); i++ {

		j = i
		t = s[i]

		for j > 0 && t < s[j-1] {

			s[j-1], s[j] = s[j], s[j-1]

			j--
		}

		s[j] = t

	}

}
