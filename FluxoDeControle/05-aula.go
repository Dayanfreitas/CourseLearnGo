package main

import "fmt"

func main() {
	x := 0
	y := 10

	if x < 10 {
		fmt.Println(x)
	}

	if !(y < 10) {
		fmt.Println(y)
	}

}
