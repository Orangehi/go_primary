//runtime包提供和go程序运行时的互操作。
package main

import (
	"fmt"
	"runtime"
)

//func GOROOT() string 返回go的root目录
//func NumCPU() int 返回本地机器的逻辑CPU数量
//func GOMAXPROCS(n int) int 设置可同时执行的CPU数
//func GC() 进行一次垃圾回收，当前一次执行垃圾回收所剩下的数据以及当前不用的数据之和达到一个百分比，go将执行一次垃圾回收

//func NumGoroutine() int返回当前存在的go程数

//func Goexit() Goexit终止调用它的go程。其它go程不会受影响。Goexit会在终止该go程前执行所有defer的函数。

//func Gosched() 使当前go程放弃处理器，以让其他go程活动。它不会挂起当前go程，因此当前go程未来会恢复

func main() {

	//	runtime.GOMAXPROCS(runtime.NumCPU()) //设置当前可以同时执行的CPU数量为本机的逻辑CPU的数量

	//fmt.Println("end")

	//fmt.Println("本地机器的逻辑CPU数量：", runtime.NumCPU())

	fmt.Println("go的root目录：", runtime.GOROOT())

}

//runtime.GOMAXPROCS(x int) x的值总在1-256的范围内

//func Gosched() 该函数的作用是暂停调用他的go程，放入带执行队列中，不久就会被调用
