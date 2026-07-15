package main

import (
	"fmt"
)

func CubePipeline(input <-chan uint8) <-chan float64 {
	output := make(chan float64)

	go func() {
		defer close(output)

		for value := range input {
			number := float64(value)
			output <- number * number * number
		}
	}()

	return output
}
func main() {

	input := make(chan uint8)

	output := CubePipeline(input)

	go func() {
		defer close(input)

		numbers := []uint8{2, 3, 4, 5, 10}

		for _, number := range numbers {
			input <- number
		}
	}()

	fmt.Println("Результат:")

	for value := range output {
		fmt.Println(value)
	}
}
