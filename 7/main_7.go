package main

import (
	"fmt"
	"sync"
)

func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for value := range c {
				out <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		defer close(ch2)
		ch2 <- 4
		ch2 <- 5
	}()

	go func() {
		defer close(ch3)
		ch3 <- 6
		ch3 <- 7
	}()

	result := FanIn(ch1, ch2, ch3)

	fmt.Println("Результат слияния:")
	for value := range result {
		fmt.Println(value)
	}
}
