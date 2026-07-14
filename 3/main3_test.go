package main

import "testing"

func TestAdd(t *testing.T) {
	m := NewStringIntMap()

	m.Add("Tanchik", 1)

	value, ok := m.Get("Tanchik")

	if !ok {
		t.Errorf("ожидался существующий ключ")
	}

	if value != 1 {
		t.Errorf("ожидалось значение 1, получено %d", value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()

	m.Add("Alina", 100)
	m.Remove("Alina")

	if m.Exists("Alina") {
		t.Errorf("ключ должен быть удален")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()

	m.Add("Kolya", 10)

	if !m.Exists("Kolya") {
		t.Errorf("ключ должен существовать")
	}

	if m.Exists("Anna") {
		t.Errorf("ключ не должен существовать")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()

	m.Add("Lena", 25)

	value, ok := m.Get("Lena")

	if !ok {
		t.Errorf("ожидался успешный поиск")
	}

	if value != 25 {
		t.Errorf("ожидалось 25, получено %d", value)
	}

	_, ok = m.Get("unknown")

	if ok {
		t.Errorf("ключ не должен существовать")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()

	m.Add("a", 1)
	m.Add("b", 2)

	copyMap := m.Copy()

	if len(copyMap) != 2 {
		t.Errorf("ожидалось 2 элемента, получено %d", len(copyMap))
	}

	if copyMap["a"] != 1 || copyMap["b"] != 2 {
		t.Errorf("копия содержит неверные значения")
	}

	copyMap["a"] = 100

	value, _ := m.Get("a")

	if value != 1 {
		t.Errorf("изменение копии не должно менять оригинал")
	}
}
