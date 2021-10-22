package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func1()
	go func2()
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 1", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 2", i)
	}
	wg.Done()
}
