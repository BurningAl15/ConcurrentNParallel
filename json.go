package main

import (
	"encoding/json"
	"fmt"
)

type Thing struct {
	Name string
	Favo int
}

type Thong struct {
	Name string `json:"name"`
	Favo int    `json:"favo"`
}

func main() {
	t1 := Thing{"Joe", 123}
	t2 := Thong{"Jack", 99}

	b1, _ := json.Marshal(t1)
	b2, _ := json.Marshal(t2)

	fmt.Println(string(b1))
	fmt.Println(string(b2))

	b3 := []byte(`{"name": "Rose", "favo": 13}`)
	t3 := &Thong{}

	if err := json.Unmarshal(b3, t3); err != nil {
		panic(err)
	}
	fmt.Println(*t3)
}
