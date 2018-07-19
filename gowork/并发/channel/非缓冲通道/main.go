//非缓冲通道
package main

import (
	"fmt"
)

//向通道发送一个数据，直到该通道被接收一个数据为止，当前Goroutine会一直被阻塞
//向通道接收一个数据，直到该通道被发送一个数据为止，当前Goroutine会被一直阻塞

//针对非缓冲通道的接收操作会在发送操作之前完成

func main() {

	c := make(chan int)

	go func() {

		fmt.Println("go")

		<-c

		fmt.Println("go end")

	}()

	fmt.Println("main")

	c <- 9

	fmt.Println("main end")

}
