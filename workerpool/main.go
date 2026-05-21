package main

import (
	"fmt"
	"sync"
)

func calculate(wg *sync.WaitGroup, input <-chan int, id int) {
	defer wg.Done()

	for x := range input {
		res := 2*x + 1
		fmt.Printf("worker %v calculated the result: %v\n", id, res)
	}
}

func main() {
	workersCnt := 3

	var wg sync.WaitGroup
	wg.Add(workersCnt)

	input := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
	}()

	for i := 0; i < workersCnt; i++ {
		go calculate(&wg, input, i)
	}

	wg.Wait()
}
