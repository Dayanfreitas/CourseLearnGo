package main

import "fmt"

func main() {
	// x := 10
	// y := returnFunc()
	// fmt.Printf("%v", y(x))
	fmt.Println(returnFunc()(10))

}

func returnFunc() func(x int) int {
	return func(x int) int {
		return (x * 10)
	}
}
