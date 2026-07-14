package main

import (
	"fmt"
)

func Correspondence(a, b []int) (bool, []int) {
	values := make(map[int]bool)

	for _, num := range a {
		values[num] = true
	}

	result := make([]int, 0)

	for _, num := range b {

		if values[num] {
			result = append(result, num)
		}
	}

	return len(result) > 0, result
}
func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	exists, result := Correspondence(a, b)

	fmt.Println("Первый слайс: ", a)
	fmt.Println("Второй слайс: ", b)
	fmt.Println("Есть ли одинаковые элементы в слайсах? ", exists)
	fmt.Println("Элементы, которые есть в обоих слайсах: ", result)

}
