package main

import (
	"fmt"
)

func Difference(slice1 []string, slice2 []string) []string {
	exists := make(map[string]bool)

	for _, item := range slice2 {
		exists[item] = true
	}

	result := make([]string, 0)

	for _, item := range slice1 {
		if !exists[item] {
			result = append(result, item)
		}
	}

	return result
}
func main() {

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := Difference(slice1, slice2)

	fmt.Println("Первый слайс: ", slice1)
	fmt.Println("Второй слайс: ", slice1)

	fmt.Println("Элементы, которых нет во втором слайсе: ", result)
}
