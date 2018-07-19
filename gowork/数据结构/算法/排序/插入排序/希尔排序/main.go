//Sell's sort 希尔排序
package main

import (
	"fmt"
)

//希尔排序是插入排序的改进版本。
//希尔排序原理：获取一个增量d1,d1小于数组的长度，一般去n/2(数组的长度),
//将数组的起始位置与数组的第（0+d1）个位置的数用插入排序
//依次对1+d1个数进行排序
//然后将增量变为d1/2,获得增量d2，把增量变为d2，重复使用希尔排序，
//一直到增量等于1。
//当增量变为1，并且已经完成了希尔排序，这时的数组就变成了一个有序数组

//希尔排序
func ShellSort(s []int) {

	lenght := len(s) / 2

	for lenght >= 1 {

		fmt.Println(lenght)
		Sort(s, lenght)
		fmt.Println(s)
		lenght = lenght / 2

	}

}

//希尔排序
func Sort(s []int, l int) {

	for i := 0; i < l; i++ { //把数组按照增树来分组,分为l组

		for j := i + l; j < len(s); j += l { //获取每一组的每一个数据

			t := s[j]
			x := j

			for x-l >= 0 && t < s[x-l] {

				s[x], s[x-l] = s[x-l], s[x]

				x -= l

			}

		}

	}

}

func main() {
	s := []int{4, 5, 3, 6, 1, 7, 9, 2, 8, 10}
	fmt.Println(s)
	ShellSort(s)
	//fmt.Println(s)
}
