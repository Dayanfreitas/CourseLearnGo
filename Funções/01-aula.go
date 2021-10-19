package main

import "fmt"

func main() {
	// basic()
	arg("manhã")
	arg("tarde")

	valor := soma(10, 10)
	fmt.Println(valor)

	total, size := somaList(10, 10, 10)
	fmt.Println(total, size)
}

func basic() {
	fmt.Println("Oi bom dia !")
}
func arg(s string) {
	if s == "manhã" {
		fmt.Println("Oi bom dia !")

	} else if s == "tarde" {
		fmt.Println("Oi boa tarde !")
	}
}

func soma(x, y int) int {
	return x + y
}

func somaList(x ...int) (int, int) {

	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x)
}
