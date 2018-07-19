//当遇到一般错误时，可以通过返回error，而当遇到不可恢复情况时，就要用到panic
package main

//为了使的编程人员可以报告程序运行期间的、且不可恢复的错误状态，go内置函数panic()
//panic常常被用作停止当前控制流程的执行，并报告一个运行时恐慌，
//panic(e interface{})接收任何类型的参数，但是比较常用的事string或者error类型

func outer() {

	inter()

}

func inter() {

	panic("panic")

}

//当函数调用panic之后，该函数就会停止执行，
//并且把流程控制返回给该函数的调用方
//运行时恐慌就会沿着调用栈层层传递，直到到达当前grountine调用栈的最顶层，
//当前goroutine调用栈的所以程序都被停止，意味着程序奔溃

func main() {
	outer()
}
