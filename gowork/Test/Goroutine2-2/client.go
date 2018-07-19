package main

import (
	"fmt"
	//	"time"
	//	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
	}
	//defer conn.Close()
	mustCopy(conn)
}
func mustCopy(conn net.Conn) {

	n, err := conn.Write([]byte("get c:/go"))

	if err != nil {

		fmt.Println("err:", err)

	} else {

		fmt.Println("n:", n)

	}

	file, err := os.Create("2.txt")
	if err != nil {

		fmt.Println("fileCreate:", err)

	}

	b1 := make([]byte, 256)

	for {

		n, err := conn.Read(b1)

		if n == 0 || err != nil {

			break

		}

		b2 := make([]byte, n)

		for i := 0; i < n; i++ {

			b2[i] = b1[i]

		}

		file.Write(b2)

		file.Close()

	}

}
