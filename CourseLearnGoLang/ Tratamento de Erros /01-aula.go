package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Isso é um erro!")

	if err != nil {
		fmt.Println(err)
	}
}
