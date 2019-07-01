package main

import "fmt"

func Y(id int) {
	fmt.Printf("%d%d%d Time %d%d%d", id, id, id, id, id, id)
	fmt.Println()
	defer fmt.Println()
	//This is printed third
	defer fmt.Printf("(%d) Mensaje X1\n", id)
	//This is printed second
	defer fmt.Printf("(%d) Mensaje X2\n", id)
	//This is printed first
	fmt.Printf("(%d) Mensaje X3\n", id)
}

func main() {
	fmt.Println()
	fmt.Println("MAAAAIN")
	fmt.Println()

	Y(1)
	defer fmt.Println("Seccion 1")
	Y(2)
	defer fmt.Println("Seccion 2")
}
