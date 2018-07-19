// Bubble sort 冒泡排序
package main

import (
	"fmt"
)

//原理：排列n-1次，分别找到每次的最大值或者最小值
//步骤：第一次排列，用第一个值与第二个值比较，第一个大于第二个就交换，再用第二个与第三个比较，一直比较到n-1个与n个，找到最大值
//第二次找到第二大或者第二小的值
//最后完成排列

func BubbleSort(s []int) {

	for i := len(s) - 1; i > 1; i-- {

		for j := 0; j < i; j++ {

			if s[j] > s[j+1] {

				s[j], s[j+1] = s[j+1], s[j]

			}

		}

	}

}

func main() {

	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8, 10}
	BubbleSort(s)
	fmt.Println(s)

}
