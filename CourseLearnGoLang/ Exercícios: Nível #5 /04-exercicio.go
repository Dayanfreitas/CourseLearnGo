// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
package main

import "fmt"

func main() {
	a := struct {
		mmap  map[string]string
		slice []string
	}{
		mmap: map[string]string{
			"name": "Dayan",
		},
		slice: []string{"sla"},
	}

	fmt.Println(a)
}
