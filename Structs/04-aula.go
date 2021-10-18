package main

import "fmt"

func main() {
	// Struct anonima
	pessoa := struct {
		name string
	}{name: "Mario"}

	fmt.Println(pessoa)
}
