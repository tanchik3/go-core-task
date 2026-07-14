package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {

	ch := Generate()

	for i := 0; i < 5; i++ {
		number := <-ch
		if number < 0 {
			t.Errorf(
				"число должно быть положительным, получено %d",
				number,
			)
		}
	}
}

func TestGenerateDifferentNumbers(t *testing.T) {

	ch := Generate()

	first := <-ch
	second := <-ch

	if first == second {
		t.Logf(
			"получены одинаковые числа: %d",
			first,
		)
	}
}

func TestGenerateChannel(t *testing.T) {
	ch := Generate()

	if ch == nil {
		t.Errorf(
			"канал не должен быть nil",
		)

	}

}
