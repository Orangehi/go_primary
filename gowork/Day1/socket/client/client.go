package main

import (
	"bytes"
	"fmt"
	"net"
)

func goclient(conn *net.TCPConn, ch chan int) {

	//defer conn.Close()

	var buffer bytes.Buffer

	b1 := make([]byte, 256)

	for {

		n, err := conn.Read(b1)

		if n == 0 || err != nil {

			break

		}

		buffer.Write(b1)

	}

	if buffer.String() != "" {

		fmt.Println(buffer.String())

	}

	ch <- 1

}

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	if err != nil {

		fmt.Println(err)

	}

	for {

		conn, err1 := net.DialTCP("tcp", nil, addr)

		if err1 != nil {

			fmt.Println(err1)

		}

		ch := make(chan int)

		go goclient(conn, ch)

		<-ch

	}

}
