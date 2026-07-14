package main

import (
	"fmt"
	"math/rand"
)

func Generate() <-chan int {

	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

func main() {
	randomNumbers := Generate()

	for i := 0; i < 5; i++ {
		number := <-randomNumbers
		fmt.Println(number)
	}
}
