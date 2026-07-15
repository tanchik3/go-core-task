package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroup struct {
	counter int
	done    chan struct{}
	mutex   sync.Mutex
}

func New() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.counter += delta

	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Done() {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.counter--

	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.done
}
func main() {

	wg := New()

	for i := 1; i <= 5; i++ {

		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)

			fmt.Println("Горутина", id, "завершилась")

		}(i)
	}

	wg.Wait()

	fmt.Println("Все горутины завершили работу")
}
