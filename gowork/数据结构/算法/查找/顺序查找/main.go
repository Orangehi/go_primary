//顺序查找，表可以为无序，或者有序。时间复杂度为O(n)
//顺序查找适合于存储结构为顺序存储或者链式存储的查找
package main

import (
	"fmt"
)

func main() {

	a := []int{2, 4, 5, 7, 8, 1, 3, 10, 3, 22, 55, 99}
	fmt.Println(a)
	value := 55
	for i := 0; i < len(a); i++ {

		if a[i] == value {

			fmt.Println(i)

		}

	}

}
