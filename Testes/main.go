package main

import "fmt"

func main() {

	fmt.Printf("Soma %v\n", Soma(1, 2, 3))
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
