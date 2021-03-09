package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err)
		}
		go proxy(conn)
		// go copyToStderr(conn)

	}
}

func proxy(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", "localhost:5678")
	if err != nil {
		fmt.Printf("Error inside proxy %v", err)
	}

	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)

}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))
		var buff [128]byte
		n, err := conn.Read(buff[:])
		if err != nil {
			fmt.Printf("finished with error %v\n", err)
			return
		}

		os.Stderr.Write(buff[:n])

	}

}
