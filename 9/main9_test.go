package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {

	input := make(chan uint8)

	output := CubePipeline(input)

	go func() {
		defer close(input)

		input <- 2
		input <- 3
		input <- 4
	}()

	var result []float64

	for value := range output {
		result = append(result, value)
	}

	expected := []float64{8, 27, 64}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestCubePipelineEmptyChannel(t *testing.T) {

	input := make(chan uint8)

	output := CubePipeline(input)

	close(input)

	var result []float64

	for value := range output {
		result = append(result, value)
	}

	if len(result) != 0 {
		t.Errorf("ожидался пустой результат")
	}
}

func TestCubePipelineOneValue(t *testing.T) {

	input := make(chan uint8)

	output := CubePipeline(input)

	go func() {
		defer close(input)

		input <- 5
	}()

	result := <-output

	expected := float64(125)

	if result != expected {
		t.Errorf("ожидалось %.0f, получено %.0f", expected, result)
	}
}
