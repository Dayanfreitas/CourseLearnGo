package main

import "fmt"

// Criando tipos
type hotdog int

var b hotdog = 10

func main() {
	x := 10

	fmt.Printf("x: %T\n", x)
	fmt.Printf("b: %T\n", b)

	x = int(b)
}
