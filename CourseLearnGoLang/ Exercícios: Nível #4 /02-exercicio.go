package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range array {
		fmt.Println(i, " ", v)
	}
	fmt.Printf("%T\n", array)

}
