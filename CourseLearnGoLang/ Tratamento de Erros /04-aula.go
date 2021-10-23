package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sb, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))
	// s := strings.NewReader("Qualquer coisa !")
	// io.Copy(f, s)
}
