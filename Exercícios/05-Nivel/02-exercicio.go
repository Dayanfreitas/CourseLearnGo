package main

import "fmt"

// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{nome: "Dayan", sobrenome: "Freitas", sabores: []string{"morango", "chocolate"}}
	pessoa2 := pessoa{nome: "Mc", sobrenome: "Braz", sabores: []string{"morango", "cheme"}}

	map_pessoas := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for _, person := range map_pessoas {
		fmt.Println(person.nome)
		fmt.Println("Sabores de sorvete")
		for _, ice_cream_flavor := range person.sabores {
			fmt.Println(ice_cream_flavor)
		}

	}
}
