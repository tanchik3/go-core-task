package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	result := FanIn(ch1, ch2)

	var values []int

	for v := range result {
		values = append(values, v)
	}

	sort.Ints(values)

	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, values)
	}
}

func TestMergeEmptyChannels(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	close(ch1)
	close(ch2)

	result := FanIn(ch1, ch2)

	var values []int

	for v := range result {
		values = append(values, v)
	}

	if len(values) != 0 {
		t.Errorf("ожидался пустой результат")
	}
}

func TestMergeOneChannel(t *testing.T) {

	ch := make(chan int)

	go func() {
		defer close(ch)

		ch <- 10
		ch <- 20
	}()

	result := FanIn(ch)

	var values []int

	for v := range result {
		values = append(values, v)
	}

	expected := []int{10, 20}

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, values)
	}
}
