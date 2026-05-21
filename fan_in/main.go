package main

import "fmt"

func collectNumbers(input1, input2 <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
			case x := <-input1:
				output <- x
			case x := <-input2:
				output <- x
			}
		}
	}()
	return output
}

func main() {
	input1 := make(chan int)
	input2 := make(chan int)

	go func() {
		for i := 0; i < 30; i++ {
			input1 <- i
		}
	}()

	go func() {
		for i := 50; i < 80; i++ {
			input2 <- i
		}
	}()

	output := collectNumbers(input1, input2)
	for i := range output {
		fmt.Println(i)
	}
}
