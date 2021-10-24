package main

// Sobre o pacote fmt

import "fmt"

func main() {
	// interpert literal
	x := "oi bom dia\ncomo vai?\t\"espero\" que tudo bem !\n"
	// RAW literal
	y := `oi bom dia\ncomo vai?\t\"espero\" que tudo bem !\n`

	fmt.Print(x)
	fmt.Print(y)
	fmt.Println()

	// Sprint
	x = "oi"
	y = "tudo bem"
	fmt.Println(fmt.Sprint(x, " ", y))

}
