package main

import (
	"fmt"
	"sync"
	"time"
)

func calculate(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	defer wg.Done()

	for x := range input {
		res := 2*x + 1
		time.Sleep(time.Second * 1)
		output <- res
	}
}

func main() {
	workersCnt := 3

	var wg sync.WaitGroup
	wg.Add(workersCnt)

	input := make(chan int)
	output := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
	}()

	go func() {
		for i := range output {
			fmt.Println("output", i)
		}
	}()

	for i := 0; i < workersCnt; i++ {
		go calculate(&wg, input, output)
	}

	wg.Wait()
	close(output)
}
