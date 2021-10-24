package main

import "fmt"

func main() {
	fmt.Print(fatorial(0))
}

func fatorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * fatorial(x-1)
}
