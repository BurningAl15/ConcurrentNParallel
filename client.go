package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//Connects to the address
	con, _ := net.Dial("tcp", "localhost:8000")
	//Close conection after all operations
	defer con.Close()
	//r is a reader from the connection
	r := bufio.NewReader(con)

	//To get the errors
	gin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Msg: ")
		//We read info from the connection but also check for errors
		msg, _ := gin.ReadString('\n')
		//We print the message from the connection
		fmt.Fprint(con, msg)
		//We read info from the connection
		resp, _ := r.ReadString('\n')
		//We print the message from the connection
		fmt.Print("Resp:", resp)
		if len(msg) == 0 || msg[0] == 'x' || msg[0] == 'X' {
			break
		}
	}
}
