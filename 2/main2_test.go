package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result := sliceExample(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}

	result := addElements(input, 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}

	copied := copySlice(input)

	input[0] = 100

	if copied[0] == 100 {
		t.Error("копия изменилась после изменения оригинала")
	}

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, copied)
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}

	result := removeElement(input, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	input := []int{1, 2, 3}

	result := removeElement(input, 10)

	if !reflect.DeepEqual(result, input) {
		t.Error("при неверном индексе слайс не должен изменяться")
	}
}
