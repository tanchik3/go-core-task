package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestDifferenceEmptySecondSlice(t *testing.T) {

	slice1 := []string{"one", "two", "three"}
	slice2 := []string{}
	expected := []string{"one", "two", "three"}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestDifferenceNoMatches(t *testing.T) {

	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e"}
	expected := []string{"a", "b", "c"}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}

func TestDifferenceAllExist(t *testing.T) {

	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b", "c"}
	expected := []string{}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"ожидалось %v, получено %v",
			expected,
			result,
		)
	}
}
