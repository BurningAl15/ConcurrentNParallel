package main

import (
	"fmt"
	"net"
)

func main() {
	con2, _ := net.Dial("tcp", fmt.Sprintf("localhost:8000"))
	defer con2.Close()
	fmt.Fprintln(con2, "KHDHÃ‘SLKJGSÃ‘LFKHGÃ‘SKFJHGKJDFHGJKLDHGLKJDFgu!")

}
