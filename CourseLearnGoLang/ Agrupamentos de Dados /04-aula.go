package main

import "fmt"

func main() {
	sabores := []string{"peperoni", "mozzarella", "abacaxi", "quatro queijos", "marguerita"}

	fmt.Println(sabores[0:2])
	fmt.Println(sabores[:])

	for i, valor := range sabores {
		fmt.Println(i, " ", valor)
	}
}
