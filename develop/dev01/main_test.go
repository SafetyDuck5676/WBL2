package main

import (
	"os"
	"testing"

	"github.com/beevik/ntp"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestPrintExactTime(t *testing.T) {
	_, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Обработка ошибки: вывод в STDERR, возврат ненулевого кода выхода в OS
		t.Errorf("Ошибка получения времени от сервера NTP: %s", err)
	}
}
