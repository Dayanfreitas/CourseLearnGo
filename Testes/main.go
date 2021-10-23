package main

import (
	"fmt"
)

func main() {

	fmt.Printf("Soma %v\n", soma(1, 2, 3))
	fmt.Printf("Multiplica %v\n", multiplica(10, 10))
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total
}
