package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func multiply(ctx context.Context, cancel context.CancelFunc, multiplier int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range input {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Millisecond * 200)
			res := val * multiplier

			if res == 10 {
				cancel()
				return
			}

			select {
			case output <- res:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan int)

	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			select {
			case input <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	output := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go multiply(ctx, cancel, i, input, output, &wg)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for res := range output {
		fmt.Printf("result: %v\n", res)
	}

	fmt.Println("done")
}
