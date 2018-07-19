//二元选择排序
package main

import (
	"fmt"
)

//原理：一次选择排序，找到最大值与最小值
//步骤：首先找到最小值和最大值，完成交换，接着找第二大第二小的值，进行交换

//返回当前数组的最大值最小值的索引
func IndexMinMax(s []int, begin int, end int) (index_Min int, index_Max int) {

	index_Min = begin
	index_Max = begin

	for i := begin; i < end+1; i++ {

		if s[index_Min] > s[i] {

			index_Min = i

		}

	}

	for i := begin; i < end+1; i++ {

		if s[index_Max] < s[i] {

			index_Max = i

		}

	}

	return

}

func TwoSelectSort(s []int) {
	end := len(s) - 1
	if len(s)%2 == 0 {

		for i := 0; i < len(s)/2; i++ {

			x, y := IndexMinMax(s, i, end)
			fmt.Println(i, end)

			if (end - i) > 1 {

				s[i], s[x] = s[x], s[i]
				s[end], s[y] = s[y], s[end]
				fmt.Println(s)

			} else {

				if s[end] < s[i] {

					s[i], s[end] = s[end], s[i]

				}

			}
			end--

		}

	} else {

		for i := 0; i < len(s)/2; i++ {

			x, y := IndexMinMax(s, i, end)

			if end-i == 2 {

				fmt.Println(x, y)
				fmt.Println(i, end)
				s[i], s[x] = s[x], s[i]
				x, y := IndexMinMax(s, i+1, end)
				temp1 := s[x]
				temp2 := s[y]
				s[i+1] = temp1
				s[i+2] = temp2

			} else {

				s[i], s[x] = s[x], s[i]
				s[end], s[y] = s[y], s[end]
				fmt.Println(s)
			}

			end--

		}

	}

}

func main() {

	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8, 10}
	TwoSelectSort(s)
	fmt.Println(s)

}
