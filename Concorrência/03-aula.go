package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Condição de corrida

func main() {
	fmt.Println("CPUS", runtime.NumCPU())

	sizeGoRutines := 1000
	contador := 0

	var wg sync.WaitGroup
	wg.Add(sizeGoRutines)
	var mu sync.Mutex

	for i := 0; i < sizeGoRutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			//Yield
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Routines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println(contador)
}
