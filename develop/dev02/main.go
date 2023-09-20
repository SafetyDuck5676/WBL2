package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result strings.Builder
	runes := []rune(s)

	var repeatCount int
	for i := 0; i < len(runes); i++ {

		if unicode.IsDigit(runes[i]) {
			digit, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", fmt.Errorf("некорректное число повторений: %w", err)
			}

			repeatCount = repeatCount*10 + digit
			if repeatCount > 0 {
				result.WriteString(strings.Repeat(string(runes[i-1]), repeatCount-1))
				repeatCount = 0
			}
		} else {
			result.WriteString(string(runes[i]))
		}
	}

	return result.String(), nil
}

func main() {
	str := "a4bc2d5e"
	unpacked, err := unpackString(str)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Распакованная строка:", unpacked)
}
