package main

import "fmt"

// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{nome: "Dayan", sobrenome: "Freitas", sabores: []string{"morango", "chocolate"}}
	pessoa2 := pessoa{nome: "Dayan", sobrenome: "Freitas", sabores: []string{"morango", "cheme"}}

	fmt.Println(pessoa1.nome)

	for i, v := range pessoa1.sabores {
		fmt.Println(i, " ", v)
	}

	fmt.Println(pessoa2.nome)
	for i, v := range pessoa2.sabores {
		fmt.Println(i, " ", v)
	}
}
