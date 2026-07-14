package main

import (
	"strings"
	"testing"
)

func TestGetType(t *testing.T) {

	if GetType(10) != "int" {
		t.Errorf("ожидался int")
	}

	if GetType(3.14) != "float64" {
		t.Errorf("ожидался float64")
	}

	if GetType("Go") != "string" {
		t.Errorf("ожидался string")
	}

	if GetType(true) != "bool" {
		t.Errorf("ожидался bool")
	}

	if GetType(complex64(1+2i)) != "complex64" {
		t.Errorf("ожидался complex64")
	}
}

func TestConvertToString(t *testing.T) {

	result := ConvertToString(42, "Go", true)

	if !strings.Contains(result, "42") {
		t.Error("нет числа")
	}

	if !strings.Contains(result, "Go") {
		t.Error("нет строки")
	}

	if !strings.Contains(result, "true") {
		t.Error("нет bool")
	}
}

func TestStringToRunes(t *testing.T) {

	runes := StringToRunes("Go")

	if len(runes) != 2 {
		t.Error("неверная длина")
	}

	if runes[0] != 'G' {
		t.Error("неверная первая руна")
	}
}

func TestHashRunes(t *testing.T) {

	runes := []rune("Hello")

	hash1 := HashRunes(runes, "go-2024")
	hash2 := HashRunes(runes, "go-2024")

	if hash1 != hash2 {
		t.Error("хэш должен быть одинаковым")
	}

	if len(hash1) != 64 {
		t.Error("SHA256 должен иметь длину 64 символа")
	}
}
