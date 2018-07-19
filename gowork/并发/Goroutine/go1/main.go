//一个go语句代表一个函数或者方法的并发执行
package main

//当go语句被执行，其中的go函数或者方法会被放到一个单独的Goroutine中
//该go函数或者方法的执行会独立于当前的Goroutine的运行
//一般情况下，当前Goroutine下，某条go语句的后面的语句并不会等到go函数执行后才执行，
//有时候，go函数后面的语句已经执行了，该go函数才刚刚执行

//go函数可以有结果申明，但是该go函数已经独立于当前Goroutine的运行，
//也就是说该go函数的返回值将没有任何意义。

//go语句和defer语句很相似，如果go后面的函数或者方法带有参数，他会对参数先求值，最后程序退出的时候才会执行函数体

func main() {

	name := "tom"

	go func(s string) {

		fmt.Println(s)

	}(name)

	//	go fmt.Println("go2")
	//	go fmt.Println("go3")
	//	go fmt.Println("go4")
	name = "jack"
	runtime.Gosched()

}
