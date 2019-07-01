package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("No se pudo iniciar serv.")
		return
	}
	defer ln.Close()
	fmt.Println("Servidor listo y escuchando.")

	con, err := ln.Accept()
	if err != nil {
		fmt.Println("No se pudo conectar.")
		return
	}
	
	defer con.Close()
	fmt.Println("Conexión establecida.")
	
	r := bufio.NewReader(con)
	msg, err := r.ReadString('\n')
	
	if err != nil {
		fmt.Println("No se pudo leer mensaje.")
	} else {
		fmt.Println("Recibido: ", msg)
	}
}
