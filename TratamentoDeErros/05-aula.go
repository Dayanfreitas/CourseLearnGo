package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Read() {
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

}

func main() {
	if _, err := os.Stat("arquivo.txt"); os.IsNotExist(err) {
		f, err := os.Create("arquivo.txt")

		if err != nil {
			panic(err)
		}

		defer f.Close()

		s := strings.NewReader("Qualquer coisa !")
		io.Copy(f, s)
		Read()

		return
	}

	Read()

}
