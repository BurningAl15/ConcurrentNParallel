package main

import (
	"fmt"
	"net"
)

func main() {
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	fmt.Fprint(con, "MSG!")
	fmt.Println("MESSAGE SENT!")
}
