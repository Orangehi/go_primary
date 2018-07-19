package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func sehandle(conn *net.TCPConn) {

	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
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

		sehandle(conn)

	}
}
