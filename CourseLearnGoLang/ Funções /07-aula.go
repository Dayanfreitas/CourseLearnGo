package main

import "fmt"

func main() {
	x := 10

	a := func(x int) {
		fmt.Println(x * 10)
	}

	a(x)
}
