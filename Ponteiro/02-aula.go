package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Valor inical:", x)
	valor(x)
	fmt.Println(x)

	pointer(&x)
	fmt.Println(x)

}

func valor(x int) {
	x++
	fmt.Println(x)

}

func pointer(x *int) {
	*x++
	fmt.Println(*x)

}
