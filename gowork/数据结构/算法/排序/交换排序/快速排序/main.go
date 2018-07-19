//Quick Sort
package main

//原理：获取一个比较元素，一般是第一个元素。将小于他的放到他的左边，大于他的放到他的右边。在对左边和右边进行快速排序
//步骤：首先取出第一个元素的值，

import (
	"fmt"
)

func QuickSort(a []int, head int, tail int) {

	key := a[head]

	p := head

	i, j := head, tail

	for i < j {

		for j > i && a[j] > key { //j指向末尾的指针，当末尾的值大于key时，j指向前一个数，即j--
			j--
		}

		if j > p { //上面跳出循环，说明末尾出现小于key的值，而j的值是该小于key值的指针，奖他们交换

			a[j], a[p] = a[p], a[j]
			p = j
		}

		for i < j && a[i] < key { //i指向前面的值的指针,当前面的值小于key时，i指向下一个值，即i++

			i++

		}

		if i < p { //上面跳出循环，说明前面出现大于key的值，当前的i即为当前值的指针，将他们交换

			a[i], a[p] = a[p], a[i]
			p = i
		}

	}

	a[p] = key

	if (p - 1) > 1 { //当p左边还有元素时，递归调用快速排序

		QuickSort(a, 0, p-1)

	}

	if (tail - p) > 1 { //当p右边还有元素时，递归调用快速排序

		QuickSort(a, p+1, tail)

	}

}

func quickSort(s []int) {

	QuickSort(s, 0, (len(s) - 1))

}

func main() {

	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8, 10}
	quickSort(s)
	fmt.Println(s)

}
