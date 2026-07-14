package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {

	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{64, 3}

	exists, result := Correspondence(a, b)

	if !exists {
		t.Errorf("ожидалось true, получено false")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestIntersectionWithoutMatches(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	expected := []int{}

	exists, result := Correspondence(a, b)

	if exists {
		t.Errorf("ожидалось false, получено true")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestIntersectionEmptySlice(t *testing.T) {

	a := []int{}
	b := []int{1, 2, 3}
	expected := []int{}

	exists, result := Correspondence(a, b)

	if exists {
		t.Errorf("ожидалось false")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestIntersectionSameValues(t *testing.T) {

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	expected := []int{1, 2, 3}

	exists, result := Correspondence(a, b)

	if !exists {
		t.Errorf("ожидалось true")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}
