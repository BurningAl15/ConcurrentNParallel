package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strings"
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

type Block struct {
	Index     int    `json:"Index"`
	Timestamp string `json:"Timestamp"`
	Hash      string `json:"Hash"`
	PrevHash  string `json:"PrevHash"`
	Payload   string `json:"Payload"`
}

type Ledger struct {
	Nombre string
}

var blockchain []Block

func Encoder() {

	block := Block{
		Index:     10,
		Timestamp: "11-05-1992",
		Hash:      "312",
		PrevHash:  "123",
		Payload:   "1231",
	}
	b, _ := json.Marshal(block)
	s := string(b)
	fmt.Println(s)
}

func Decoder(_msg string) {

	bytes := []byte(_msg)
	var block_ Block

	json.Unmarshal(bytes, &block_)
	blockchain = append(blockchain, block_)
	for l := range blockchain {
		fmt.Printf("-------------------\n")
		fmt.Printf("Data File Number %d\n", l)
		fmt.Printf("-------------------\n")
		fmt.Printf("Index=%d\n", blockchain[l].Index)
		fmt.Printf("Timestamp=%s\n", blockchain[l].Timestamp)
		fmt.Printf("Hash=%s\n", blockchain[l].Hash)
		fmt.Printf("Prevhash=%s\n", blockchain[l].PrevHash)
		fmt.Printf("Payload=%s\n", blockchain[l].Payload)
		fmt.Printf("-------------------\n")
		fmt.Println(" ")
	}
}

const (
	PROTOCOL  = "tcp"
	LOCALHOST = "10.11.98.229:8000"
)

func send(n string) {
	host := HOSTS[rand.Intn(len(HOSTS))]
	resp := fmt.Sprintf("%s", n)
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

	send(msg)
}

func start() {
	var num string
	fmt.Print("Mensaje: ")
	fmt.Scanf("%s\n", &num)
	send(num)
}

func main() {
	Decoder(`{"Index":0,"Timestamp":"11-05-1992","Hash":"312","Prevhash":"123","Payload":"1231"}`)
	Decoder(`{"Index":1,"Timestamp":"11-05-1992","Hash":"312","Prevhash":"123","Payload":"1231"}`)
	Decoder(`{"Index":2,"Timestamp":"11-05-1992","Hash":"312","Prevhash":"123","Payload":"1231"}`)

	fmt.Println("Elements in blockchain: ", len(blockchain))
}
