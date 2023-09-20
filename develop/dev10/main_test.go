package main

import (
	"os"
	"testing"
)

func TestHandleError(t *testing.T) {
	// Установка переменной окружения EXITCODE в пустое значение перед тестом
	os.Setenv("EXITCODE", "")

	// Тестирование обработки ошибки, которая не является ошибкой EOF
	var err error

	// Установка переменной окружения EXITCODE в "0" перед вызовом handleError
	os.Setenv("EXITCODE", "0")
	handleError(err)

	// Ожидается, что программа завершится с кодом 1, так как handleError изменит значение EXITCODE на "1"
	if got := os.Getenv("EXITCODE"); got != "0" {
		t.Errorf("Expected EXITCODE=0, got %s", got)
	}
}
