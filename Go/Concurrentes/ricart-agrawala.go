package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	fmt.Println("Esperando conexion...")
	conn, _ := ln.Accept()
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Print(msg)

	conn.Close()
	ln.Close()
}
