package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)

	for key, value := range m.data {
		copyMap[key] = value
	}

	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}
func main() {
	data := NewStringIntMap()

	data.Add("Ivan", 10)
	data.Add("Masha", 20)

	fmt.Println("Мапа с данными: ", data.Copy())

	if data.Exists("Tanya") {
		fmt.Println("Ключ существует")
	} else {
		fmt.Println("Ключ не найден")
	}

	value, ok := data.Get("Ivan")
	if ok {
		fmt.Println("Значение по ключу:", value)
	} else {
		fmt.Println("Ключ Ivan не найден")
	}

	data.Remove("Ivan")
	fmt.Println("Мапа после удаления элемента:", data.Copy())

	copyMap := data.Copy()
	fmt.Println("Копия карты:", copyMap)
	copyMap["Dasha"] = 32
	fmt.Printf("Измененная копия: %v\nОригинал после изменения копии: %v\n", copyMap, data.Copy())
}
