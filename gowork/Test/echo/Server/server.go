package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {

	listen, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("listen:", err)
	}

	for {

		conn, err1 := listen.Accept()
		if err1 != nil {

			//fmt.Println("accept:", err1)
			continue

		}

		go handleConn(conn)

	}

}
