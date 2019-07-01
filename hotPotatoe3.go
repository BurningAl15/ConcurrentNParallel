package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var HOSTS = []string{"10.11.97.199:8000",
	"10.11.98.201:8000",
	"10.11.98.213:8000",
	"10.11.98.211:8000",
	"10.11.98.210:8000",
	"10.11.98.215:8000",
	"10.11.98.214:8000",
	"10.11.98.226:8000",
	"10.11.97.225:8000",
	"10.11.97.218:8000",
	"10.11.97.219:8000",
	"10.11.98.207:8000",
	"10.11.98.205:8000",
	"10.11.98.229:8000"}

const (
	PROTOCOL  = "tcp"
	LOCALHOST = "10.11.98.229:8000"
)

func send(n int) {
	host := HOSTS[rand.Intn(len(HOSTS))]
	resp := fmt.Sprintf("%d", n)
	con, _ := net.Dial(PROTOCOL, host)
	fmt.Println(n)
	fmt.Println(host)
	defer con.Close()
	fmt.Fprintln(con, resp)
}

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	msg = strings.TrimSpace(msg)
	if n, err := strconv.Atoi(msg); err == nil {
		if n <= 0 {
			fmt.Println("Me tocÃ³ perder!")
		} else {
			send(n - 1)
		}
	}
}

func start() {
	var num int
	fmt.Print("Numerito: ")
	fmt.Scanf("%d\n", &num)
	send(num)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ln, _ := net.Listen(PROTOCOL, LOCALHOST)
	defer ln.Close()
	go start()
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}
