package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// file, err := os.Stat("./arquivo.txt")
	// // os.IsNotExist(err) {
	// // 	// your file does not exist
	// // 	fmt.Println(err)
	// // 	return

	// // }
	// fmt.Println(file)

	f, err := os.Create("arquivo.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	s := strings.NewReader("Qualquer coisa !")
	io.Copy(f, s)
}
