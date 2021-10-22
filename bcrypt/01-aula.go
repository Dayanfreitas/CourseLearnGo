package main

import (
	"fmt"
	//"golang.org/x/crypto/bcrypt"
)

func main() {
	// fmt.Print()
	password := "21novembro1"
	// GenerateFromPassword(password []byte, cost int) ([]byte, error)
	sb, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sb)
	fmt.Println(string(sb))
}
