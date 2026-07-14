package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(slice []int) []int {
	honest := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			honest = append(honest, num)
		}
	}
	return honest
}

func addElements(slice []int, n int) []int {
	return append(slice, n)
}

func copySlice(slice []int) []int {
	copied := make([]int, len(slice))
	copy(copied, slice)
	return copied
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)

	return result
}

func main() {

	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) // числа от 0 до 99
	}

	fmt.Println("Изначальный слайс: ", originalSlice)
	fmt.Println("Слайс с четными элементами: ", sliceExample(originalSlice))
	fmt.Println("Слайс с доп. элементом: ", addElements(originalSlice, 12))
	fmt.Println("Копия слайса: ", copySlice(originalSlice))
	copiedSlice := copySlice(originalSlice)
	originalSlice[1] = 56
	fmt.Printf("Сравнение слайсов после изменения:\n Оригинал: %v\n Копия: %v\n", originalSlice, copiedSlice)
	fmt.Println("Слайс без элемента по индексу 2: ", removeElement(originalSlice, 2))
}
