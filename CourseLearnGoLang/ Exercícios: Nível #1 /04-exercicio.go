package main

import "fmt"

// - Crie um tipo. O tipo subjacente deve ser int.
// - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// - Na função main:
//     1. Demonstre o valor da variável "x"
//     2. Demonstre o tipo da variável "x"
//     3. Atribua 42 à variável "x" utilizando o operador "="
//     4. Demonstre o valor da variável "x"
type hotdog int

var x hotdog

func main() {
	fmt.Println("Exercícios 04")
	fmt.Printf("X: %v %T", x, x)

	x = 42

	fmt.Println(x)
}
