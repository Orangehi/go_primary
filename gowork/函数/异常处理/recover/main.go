//recover为go的内置函数，可以使的当前程序从运行时恐慌中恢复并重新获得流程控制权
package main

//"fmt"

//如果当前程序出现运行时恐慌，调用recover()将会返回一个 interface{}类型的值

//当程序出现运行时恐慌，就会停止运行，所以调用recover并没有作用。
//这时候就需要配套使用defer函数
//defer func(){

//	if err := recover();err != nil{

//		fmt.Println(err)

//	}

//}()

func main() {

}
