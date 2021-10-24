package main

import (
	"fmt"
	"sort"
)

// type carro

// type Interface interface {
// 	// Len is the number of elements in the collection.
// 	Len() int
// 	// Less reports whether the element with
// 	// index i should sort before the element with index j.
// 	Less(i, j int) bool
// 	// Swap swaps the elements with indexes i and j.
// 	Swap(i, j int)
// }

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// type ordenarPorConsumo []carro
// type ordenarPorLucro []carro

type OrdenarPorPotencia []carro

func (a OrdenarPorPotencia) Len() int           { return len(a) }
func (a OrdenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrdenarPorPotencia) Less(i, j int) bool { return a[i].potencia < a[j].potencia }

type OrdenarPorConsumo []carro

func (a OrdenarPorConsumo) Len() int           { return len(a) }
func (a OrdenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrdenarPorConsumo) Less(i, j int) bool { return a[i].consumo > a[j].consumo }

type OrdenarPorLucro []carro

func (a OrdenarPorLucro) Len() int           { return len(a) }
func (a OrdenarPorLucro) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrdenarPorLucro) Less(i, j int) bool { return a[i].consumo < a[j].consumo }

func main() {
	cars := []carro{
		{nome: "Celta", potencia: 20, consumo: 20},
		{nome: "KA", potencia: 50, consumo: 14},
		{nome: "Porche", potencia: 1000, consumo: 5},
	}

	fmt.Println("INICIAL->", cars)

	fmt.Println("Potencia")
	sort.Sort(OrdenarPorPotencia(cars))
	fmt.Println(cars)

	fmt.Println("Consumo")
	sort.Sort(OrdenarPorConsumo(cars))
	fmt.Println(cars)

	fmt.Println("Lucro")
	sort.Sort(OrdenarPorLucro(cars))
	fmt.Println(cars)
	fmt.Println("_______")

}
