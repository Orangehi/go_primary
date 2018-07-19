//selection sort 选择排序
package main

import (
	"fmt"
)

//选择排序的原理：数组的第一个元素等于当前数组的最小值或者最大值，第二个元素等于当前数组的最小值或者最大值，依次进行
//选择排序的步骤：首先找到数组的最大值或者最小值，放到第一个位置
//接着找到数组第二小的值或者大的值放到第二位置
//依次进行
//直到数组最后一个位置

//求数组中最小值的索引
func IndexMin(s []int, begin int) int {

	k := begin

	for j := begin + 1; j < len(s); j++ {

		if s[k] > s[j] {

			k = j

		}

	}

	return k

}

//选择排序
func SelectSort(s []int) {

	for i := 0; i < len(s); i++ {

		j := i

		k := IndexMin(s, j)

		if k != i {
			s[i], s[k] = s[k], s[i]
		}

	}

}

func main() {

	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8}
	SelectSort(s)
	fmt.Println(s)

}
