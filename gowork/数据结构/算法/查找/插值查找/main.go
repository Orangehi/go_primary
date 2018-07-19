//插值查找，是二分查找的优化 时间复杂度：O(log2 (log2 n))
package main

import (
	"fmt"
)

//插值查找适用于表长数据量大，且表中的数据分布均匀，插值查找比二分查找效果更好
//当表中数据浮动比较大，相邻数据差距比较大时，插值查找将不适用于该表查找

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(InsertValueSearch(a, 41, 0, 10))

}

func InsertValueSearch(a []int, value int, begin int, end int) int {

	if begin > end {

		return -1

	}

	mid := 0
	mid = begin + (value-a[begin])/(a[end]-a[begin])*(end-begin)

	if mid == value {

		return mid

	} else if mid > value {

		return InsertValueSearch(a, value, begin, mid-1)

	} else {

		return InsertValueSearch(a, value, mid+1, end)

	}

}
