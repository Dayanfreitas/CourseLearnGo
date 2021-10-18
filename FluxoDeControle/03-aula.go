package main

import "fmt"

func main() {

	for horas := 0; horas <= 24; horas++ {
		// fmt.Printf("Hora (s): %v ", horas)

		for x := 0; x < 60; x++ {
			fmt.Printf("Hora (s): %v Minuto (s): %v", horas, x)
			fmt.Println()
			// fmt.Print(" ", x)
		}
		fmt.Println()
	}
}
