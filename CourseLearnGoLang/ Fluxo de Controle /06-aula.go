package main

import "fmt"

func main() {
	x := 0

	if x < 10 {
		fmt.Println("Menor que 10 ")
	} else if x > 100 {
		fmt.Println("Maior que cem")
	} else {
		fmt.Println("Nem um dos dois")
	}

}
