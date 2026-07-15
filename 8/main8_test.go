package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {

	wg := New()

	var counter int32

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			time.Sleep(10 * time.Millisecond)

			atomic.AddInt32(&counter, 1)

		}()
	}

	wg.Wait()

	if counter != 5 {
		t.Errorf("ожидалось 5, получено %d", counter)
	}
}

func TestWaitWithoutGoroutines(t *testing.T) {

	wg := New()
	wg.Add(0)
	wg.Wait()
}

func TestOneGoroutine(t *testing.T) {

	wg := New()

	var finished bool

	wg.Add(1)

	go func() {
		defer wg.Done()
		finished = true
	}()

	wg.Wait()

	if !finished {
		t.Error("горутина не завершилась")
	}
}
