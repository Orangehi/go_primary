package main

import (
	"fmt"
	"net"
)

func serhandle(ch chan int, conn *net.TCPConn) {

	defer conn.Close()

	//conn.SetReadDeadline(time.Now().Add(time.Minute * 3))

	var s string

	fmt.Scanf("%s", &s)

	_, err := conn.Write([]byte(s))

	//fmt.Println(n)

	if err != nil {

		fmt.Println(err)

	}

	ch <- 1

}

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")

	if err != nil {

		fmt.Println(err)

	}

	ln, err1 := net.ListenTCP("tcp", addr)

	if err1 != nil {

		fmt.Println(err1)

	}

	for {

		conn, err2 := ln.AcceptTCP()

		if err2 != nil {

			continue

		}

		ch := make(chan int)

		go serhandle(ch, conn)

		<-ch

	}

}
