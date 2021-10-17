package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	a := "e"
	b := "é"
	// c := ""

	fmt.Printf("%v %v", a, b)
	c := []byte(a)
	d := []byte(b)

	fmt.Printf("%v %v", c, d)

	e := 10
	f := 10.0
	fmt.Printf("%T %T", e, f)

}
