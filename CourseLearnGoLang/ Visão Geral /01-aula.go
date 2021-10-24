package main

import "fmt"

func main() {
	number, erro := fmt.Println("Print")
	fmt.Println(number, erro)

	//Ignora
	_, error := fmt.Println("Print")
	fmt.Println(error)
}
