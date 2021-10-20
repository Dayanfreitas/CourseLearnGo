package main

import (
	"fmt"
	"sort"
)

func main() {
	si := []int{21, 312, 123, 12, 3145}
	fmt.Println(si)

	sort.Ints(si)

	fmt.Println(si)
}
