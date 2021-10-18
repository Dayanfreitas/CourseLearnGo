package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{nome: "João", idade: 30}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 1000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	//Promovido
	fmt.Println(pessoa2.nome)
}
