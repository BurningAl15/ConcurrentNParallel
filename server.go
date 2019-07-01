package main

import (
	"bufio"
	"fmt"
	"net"
)

//In general here we handle the connection
//Then we got the message from the conection
//Finally we print the message in the server screen
//and close the conection
//Also if the message lenght is 0 or starts with 'X' or 'x', we close the
//connection instantly after we show the message
func handle(con net.Conn, id int) {
	//Close the conection before every operation is done
	defer con.Close()
	//Reads data from the connection
	r := bufio.NewReader(con)

	//While for
	for {
		//Read the string data
		msg, _ := r.ReadString('\n')
		//Prints the message from the conection
		fmt.Fprint(con, msg)
		//Prints message in screen
		fmt.Printf("Con%d: %s", id, msg)

		//If the message lenght is 0 or the first letter of the word is x
		if len(msg) == 0 || msg[0] == 'x' || msg[0] == 'X' {
			break
		}
		fmt.Printf("Con%d cerrada!\n", id)
	}
}

func main() {
	//net.Listen(network type, address)
	//ln is Listener type
	ln, _ := net.Listen("tcp", "localhost:8000")

	//We close the listener after all operations are done
	defer ln.Close()
	fmt.Println("Servidor listo y escuchando.")

	cont := 0
	for {
		//We got the connection
		//We count every message
		con, _ := ln.Accept()
		go handle(con, cont)
		cont++
	}
}
