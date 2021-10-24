package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Condição de corrida

func main() {
	fmt.Println("CPUS", runtime.NumCPU())

	sizeGoRutines := 100
	var contador int64

	var wg sync.WaitGroup
	wg.Add(sizeGoRutines)

	for i := 0; i < sizeGoRutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println("CONTADOR", atomic.LoadInt64(&contador))
			runtime.Gosched()

			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println(contador)
}
