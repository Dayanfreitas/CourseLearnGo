package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca"}

	for i, valor := range slice {
		fmt.Println(i, " ", valor)
	}
}
