package main

import "fmt"

var y = 10

func qualquerCoisa(x int) {

	fmt.Printf("Y: %v\n", y)
	fmt.Printf("X: %v\n", x)
}

func main() {
	x := 20
	qualquerCoisa(x)
}
