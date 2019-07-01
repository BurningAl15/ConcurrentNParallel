package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Fprint(con, msg)
	fmt.Printf("Recibido: %s", msg)
}
