package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"atirei", "o", "pau", "no", "gato"}
	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)
}
