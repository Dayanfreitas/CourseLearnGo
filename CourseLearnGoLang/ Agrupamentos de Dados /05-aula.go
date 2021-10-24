package main

import "fmt"

func main() {
	// slice := []string{"banana", "maçã", "jaca"}
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	slice1 = append(slice1, 4, 5, 6)

	for i, valor := range slice1 {
		fmt.Println(i, " ", valor)
	}

	slice1 = append(slice1, slice2...)
	for i, valor := range slice1 {
		fmt.Println(i, " ", valor)
	}
}
