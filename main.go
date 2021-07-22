package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9000")
	l, _ := net.ListenTCP("tcp", tcpAddr)
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(c)
	}
}

func handle(c net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		c.Close()
	}()

	// For the case if client opens the connection, but does nothing
	c.SetDeadline(time.Now().Add(5 * time.Second))
	// c.SetReadDeadline(time.Now().Add(5 * time.Second))
	// c.SetWriteDeadline(time.Now().Add(5 * time.Second))

	r := bufio.NewReader(c)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("The end!")
			break
		}
		if err != nil {
			break
		}
		response(b, c)
	}
}

func response(b []byte, c net.Conn) {
	c.Write(b)
	// panic("Don't")
}
