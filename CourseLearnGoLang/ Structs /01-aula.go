package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{nome: "João", sobrenome: "da Silva", fumante: false}
	cliente2 := cliente{"João", "Paulo", true}
	// cliente1 := cliente{nome: "João", sobrenome: "da Silva", fumante: false}
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
