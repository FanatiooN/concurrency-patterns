package main

import (
	"fmt"
	"sync"
	"time"
)

func calculate1(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	defer wg.Done()

	for x := range input {
		res := 2*x + 1
		time.Sleep(time.Second * 3)
		output <- res
	}
}

func calculate2(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	defer wg.Done()

	for x := range input {
		res := (x - 1) / 2
		time.Sleep(time.Second * 2)
		output <- res
	}
}

func main() {
	workersCnt := 3

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	input := make(chan int)
	temp := make(chan int)
	output := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
	}()

	for i := 0; i < workersCnt*2; i++ {
		wg1.Add(1)
		go calculate1(&wg1, input, temp)
	}

	for i := 0; i < workersCnt; i++ {
		wg2.Add(1)
		go calculate2(&wg2, temp, output)
	}

	go func() {
		wg1.Wait()
		close(temp)
	}()

	go func() {
		wg2.Wait()
		close(output)
	}()

	for i := range output {
		fmt.Println("output", i)
	}
}
