package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

func handleconn(conn net.Conn) {

	defer conn.Close()

	for {

		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		strReq, err := read(conn)

		if err != nil {

			if err == io.EOF {

				fmt.Println("The connection is close by another side(server)\n")

			} else {

				fmt.Printf("Read error : %s\n", err)

			}

			break

		}

		fmt.Printf("Received Request %s(server)\n", strReq)

		RespMsg := fmt.Sprintf("the cube root of %s\n", strReq)

		n, err := write(conn, RespMsg)

		if err != nil {

			fmt.Printf("write error :%s,%d\n", err, n)

		}

		fmt.Println("The Response is %s\n", RespMsg)

	}

}

func read(conn net.Conn) (string, error) {

	readBytes := make([]byte, 1)

	var buffer bytes.Buffer

	for {

		_, err := conn.Read(readBytes)

		if err != nil {

			return "", err

		}

		readByte := readBytes[0]

		if readByte == DELIMITER {

			break

		}

		buffer.WriteByte(readByte)

	}

	return buffer.String(), nil

}

func write(conn net.Conn, contant string) (int, error) {

	var buffer bytes.Buffer

	buffer.WriteString(contant)

	buffer.WriteByte(DELIMITER)

	return conn.Write(buffer.Bytes())

}

func servergo() {

	//var listener net.listener

	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)

	if err != nil {

		fmt.Printf("Listen error is:%s\n", err)

		return

	}

	defer listener.Close()

	fmt.Printf("Got Listenner for the server.(listen address:%s\n)", listener.Addr())

	for {

		conn, err := listener.Accept()

		if err != nil {

			fmt.Printf("Accept error : %s\n", err)

		}

		fmt.Printf("a connection with a client application.(remote address :%s)\n", conn.RemoteAddr())

		go handleconn(conn)

	}

}

func cilentgo(i int) {

	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)

	if err != nil {

		fmt.Printf("the client connent is file : %s", err)

		return

	}

	defer conn.Close()

	fmt.Printf("the cclient server.(remote address:%s;local address:%s)\n", conn.RemoteAddr(), conn.LocalAddr())

	time.Sleep(200 * time.Millisecond)

	resNum := 5

	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))

	for i := 0; i < resNum; i++ {

		st := "123"

		n, err := write(conn, st)

		if err != nil {

			fmt.Printf("this client write error : %s,%d\n", err, n)

			continue
		}

		fmt.Printf("the client request : %s\n", st)

	}

	for i := 0; i < resNum; i++ {

		Resp, err := read(conn)

		if err != nil {

			if err == io.EOF {

				fmt.Println("The connection is close by another side(server)\n")

			} else {

				fmt.Printf("Read error : %s\n", err)

			}

			break

		}

		fmt.Printf("the client received : %s\n", Resp)

	}

}

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go servergo()

	time.Sleep(500 * time.Millisecond)

	go cilentgo(1)

	wg.Wait()

}
