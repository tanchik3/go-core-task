package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func GetType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func ConvertToString(values ...interface{}) string {
	result := ""

	for _, value := range values {
		switch v := value.(type) {
		case int:
			result += strconv.Itoa(v)
		case float64:
			result += strconv.FormatFloat(v, 'f', -1, 64)
		case string:
			result += v
		case bool:
			result += strconv.FormatBool(v)
		case complex64:
			result += fmt.Sprintf("%v", v)
		default:
			result += fmt.Sprintf("%v", v)
		}
	}

	return result
}

func StringToRunes(s string) []rune {
	return []rune(s)
}

func HashRunes(runes []rune, salt string) string {
	middle := len(runes) / 2

	newRunes := append([]rune{}, runes[:middle]...)
	newRunes = append(newRunes, []rune(salt)...)
	newRunes = append(newRunes, runes[middle:]...)

	hash := sha256.Sum256([]byte(string(newRunes)))

	return hex.EncodeToString(hash[:])
}
func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Println("numDecimal:", GetType(numDecimal))
	fmt.Println("numOctal:", GetType(numOctal))
	fmt.Println("numHexadecimal:", GetType(numHexadecimal))
	fmt.Println("pi:", GetType(pi))
	fmt.Println("name:", GetType(name))
	fmt.Println("isActive:", GetType(isActive))
	fmt.Println("complexNum:", GetType(complexNum))

	str := ConvertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Printf("Строка: %s\n", str)

	runes := StringToRunes(str)
	fmt.Printf("Руны: %v\n", runes)

	hash := HashRunes(runes, "go-2024")
	fmt.Printf("SHA256: %v\n", hash)

}
