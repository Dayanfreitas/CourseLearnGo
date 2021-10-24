package main

//callback

import "fmt"

func main() {
	fmt.Println(returnFunc(10, mult))
}

func mult(x int) int {
	return x * 10
}

func returnFunc(x int, f func(x int) int) int {
	return f(x)
}
