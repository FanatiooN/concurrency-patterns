package main

import "fmt"

func generateNumbers(minNum, maxNum int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for i := minNum; i <= maxNum; i++ {
			output <- i
		}
	}()

	return output
}

func main() {
	generator := generateNumbers(1, 15)

	for i := range generator {
		fmt.Println(i)
	}
}
