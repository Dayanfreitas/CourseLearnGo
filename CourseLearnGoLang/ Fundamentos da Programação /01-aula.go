package main

import "fmt"

var x bool

func main() {

	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuido

	//Operadores relacionais
	y := (10 < 0)
	fmt.Println(y)
}
