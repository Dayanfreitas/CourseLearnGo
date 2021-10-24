package main

func main() {
	slice := []int{10, 20, 30, 40, 50}
	somaList(slice...)
}

func somaList(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma
}
