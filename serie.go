package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	puerto := 0
	fmt.Println("Ingrese el puerto: ")
	fmt.Scanf("%d\n", &puerto)

	ln, _ := net.Listen("tcp", fmt.Sprintf("localhost:800%d", puerto))
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Print("Recibido: ", msg)

	con2, _ := net.Dial("tcp", fmt.Sprintf("localhost:%d", (puerto+1)%5))
	defer con2.Close()
	fmt.Fprintln(con2, "KHDHÃ‘SLKJGSÃ‘LFKHGÃ‘SKFJHGKJDFHGJKLDHGLKJDFgu!")
}
