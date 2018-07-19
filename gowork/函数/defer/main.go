//defer 被用作于预定对一个函数的调用
package main

import (
	"fmt"
)

//被defer调用的函数叫作延迟函数，延迟函数只能出现在函数或方法的内部
//当函数体或者方法内部出现多个延迟函数时，延迟函数的调用应当符合先进后出，相当于栈的形式。
//延迟函数，主要用作于释放资源和异常处理。一般放在函数体开始出，或者紧跟在申请资源语句的后面

//延迟函数 defer func(func(x)x)，当出现延迟函数需要参数的时候，
//会首先对延迟函数内的参数求值。会先计算出func(x)x然后执行下面语句

func begin(st string) string {

	fmt.Println("begin:", st)
	return "end"

}

func end(st string) {

	fmt.Println("end st", st)

}

func record() {

	defer end(begin("begin"))

	fmt.Println("record")

}

//函数record（）会先求出begin（“begin”）的值，然后在打印record，最好执行延迟函数end（）
//所以如果延迟函数带有参数，会先求出该参数的值

//延迟函数的主题，也就是延迟函数的内部表达式不会被先求值。
func Pr() {

	for i := 0; i < 5; i++ {

		defer func() {

			fmt.Println(i)

		}()

	}

}

//该函数内部的延迟函数，其函数主体打印变量i的值，i的值为函数即将结束时的数据5，所以会输出5,5,5,5,5,

//如果延迟函数是一个匿名函数，且在包含延迟函数的函数中有对返回值的命名
//那么在延迟函数中对该命名的操作，将会是最后的返回值
func Back() (i int) {

	defer func() {
		i++
	}()

	i++
	return

}

func main() {

}
