package main

import (
	"fmt"
)

func main() {
	var nome int

	_, err := fmt.Scan(&nome)
	if err != nil {
		panic(err)
	}
}
