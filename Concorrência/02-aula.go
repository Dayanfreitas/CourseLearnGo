package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Condição de corrida

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())

	sizeGoRutines := 1000
	contador := 0

	var wg sync.WaitGroup
	wg.Add(sizeGoRutines)

	for i := 0; i < sizeGoRutines; i++ {
		go func() {
			v := contador
			//Yield
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Routines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Routines", runtime.NumGoroutine())
	fmt.Println(contador)
}
