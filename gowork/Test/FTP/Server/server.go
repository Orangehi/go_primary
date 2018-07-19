package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

func Diroperation(buffer bytes.Buffer, conn *net.TCPConn) {

	//	st := ""

	if n := strings.Index(buffer.String(), "cd"); n != -1 && n < 6 {

		fmt.Println("cd-", strings.TrimSpace(strings.TrimLeft(buffer.String(), "cd")))

		ndir, err1 := os.Getwd()
		if err1 != nil {
			fmt.Println("getwd err1:", err1)

		} else {
			fmt.Println("nowdir:", ndir)

		}

		err2 := os.Chdir(strings.TrimSpace(strings.TrimLeft(buffer.String(), "cd")))
		if err2 != nil {

			fmt.Println("chdir err2:", err2)

		}

		ndir1, err3 := os.Getwd()
		if err3 != nil {
			fmt.Println("getwd err2:", err3)

		} else {
			fmt.Println("after cd nowdir:", ndir1)

		}

	} else if n := strings.Index(buffer.String(), "ls"); n != -1 && n < 6 {

		fmt.Println("ls-", strings.TrimSpace(strings.TrimLeft(buffer.String(), "ls")))

		ndir, err1 := os.Getwd()
		if err1 != nil {
			fmt.Println("getwd err1:", err1)

		} else {
			fmt.Println("nowdir:", ndir)

		}

		file, err := os.Open(ndir)
		if err != nil {

			fmt.Println("fileopen:", err)

		}

		s1, err1 := file.Readdirnames(-1)

		if err1 != nil {

			fmt.Println("readdir-", err1)

		}

		for _, s := range s1 {

			fmt.Println(s)

		}

	} else if n := strings.Index(buffer.String(), "get"); n != -1 && n < 6 {

		fmt.Println("发送文件-", strings.TrimSpace(strings.TrimLeft(buffer.String(), "get")))

		file, err1 := os.Open("C:/go/go_workspace/src/Test/FTP/Server/2.txt")
		if err1 != nil {

			fmt.Println("fileopen:", err1)

		}
		b1 := make([]byte, 256)

		n1, err2 := file.Read(b1)
		if err2 != nil {

			fmt.Println("fileread:", err2)

		}

		file.Close()

		b2 := make([]byte, n1)

		for i := 0; i < n1; i++ {

			b2[i] = b1[i]

		}

		fmt.Println(b2)

		_, err3 := conn.Write(b2)
		if err3 != nil {

			fmt.Println("err3:", err3)

		}

	} else if n := strings.Index(buffer.String(), "send"); n != -1 && n < 6 {

		fmt.Println("接收文件:", strings.TrimSpace(strings.TrimLeft(buffer.String(), "send")))

		file, err := os.Create("1.txt")
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

	} else {

		fmt.Println("命令出错")

	}

}

func handleser(conn *net.TCPConn, ch chan int) {

	//defer conn.Close()

	var buffer bytes.Buffer

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

		buffer.Write(b2)

		Diroperation(buffer, conn)

	}

	ch <- 1

}

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	if err != nil {

		fmt.Println("errResolveTCPAddr:", err)

	}

	ln, err1 := net.ListenTCP("tcp", addr)
	if err1 != nil {

		fmt.Println("ListenTCP", err1)

	}

	for {

		conn, err2 := ln.AcceptTCP()
		if err2 != nil {

			continue

		}

		ch := make(chan int)

		go handleser(conn, ch)

		<-ch

	}

}
