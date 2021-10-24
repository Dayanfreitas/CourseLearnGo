package main

import "fmt"

func main() {
	x := 10
	y := 10
	z := &x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)

	fmt.Println()

	fmt.Println(&x)
	fmt.Println(z)
	fmt.Println(&y)

	fmt.Printf("%T, %T, %T", x, y, z)
}
