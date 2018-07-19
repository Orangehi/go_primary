//二分查找必须是有序的，如果为无序需得先进行排序操作
package main

import (
	"fmt"
)

//时间复杂度为O(log2 N)
//对于静态表，排序稳定，使用二分查找有一点的效率
//对于动态表，排序不稳定，每次都得先进行排序操作，不推荐使用。
//静态表：排序稳定，一次排序后不再变化
//动态表：需要进行大量的插入和删除操作，排序不稳定。

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v := 10

	//fmt.Println(BinarySearch(a, v))

	fmt.Println(BinarySearch2(a, v, 1, 10))

}

//迭代方式
func BinarySearch(a []int, value int) int {

	low, high := 0, 0
	low = a[0]
	high = a[len(a)-1]
	mid := 0
	for low < high {

		mid = (low + high) / 2
		fmt.Println("mid", mid)
		if mid == value {

			return mid

		}
		if mid < value {

			low = mid + 1

		} else if mid > value {

			high = mid - 1

		}

	}

	return -1

}

//递归方式
func BinarySearch2(a []int, value int, low int, high int) int {

	mid := (low + high) / 2

	if low > high {

		return -1
	}

	if mid == value {

		return mid

	}

	if mid < value {

		return BinarySearch2(a, value, mid+1, high)

	} else {

		return BinarySearch2(a, value, low, mid-1)

	}

}
