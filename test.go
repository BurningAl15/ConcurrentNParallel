package main

import "fmt"

func x(id int) {
	defer fmt.Printf("(%d) Mensaje X1\n", id)
	defer fmt.Printf("(%d) Mensaje X2\n", id)
	fmt.Printf("(%d) Mensaje X3\n", id)
	defer fmt.Printf("(%d) Mensaje X4\n", id)
	defer fmt.Printf("(%d) Mensaje X5\n", id)
}

func main() {
	fmt.Println("Esto es main")
	defer fmt.Println("Esto sigue siendo main (1)")
	defer x(1)
	defer fmt.Println("Esto sigue siendo main (2)")
	fmt.Println("Esto sigue siendo main (3)")
	x(2)
	defer x(3)
	defer fmt.Println("Esto sigue siendo main (3)")
	defer fmt.Println("Esto sigue siendo main (4)")
	defer fmt.Println("Esto sigue siendo main (5)")
}
