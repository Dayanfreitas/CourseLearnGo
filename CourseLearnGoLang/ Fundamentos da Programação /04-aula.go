package main

import "fmt"

func main() {
	var i uint16
	// i = 65536 -> overflow
	i = 65535
	fmt.Println(i)
	fmt.Println("####")

	i++
	fmt.Println(i)
	fmt.Println("####")

	i++
	fmt.Println(i)
	fmt.Println("####")
}
