package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Установка соединения с сервером NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Обработка ошибки: вывод в STDERR, возврат ненулевого кода выхода в OS
		log.Printf("Ошибка получения времени от сервера NTP: %s", err)
		os.Exit(1)
	}

	// Вычисление разницы времени от текущего момента до точного времени
	duration := time.Until(ntpTime)

	// Получение точного времени
	exactTime := time.Now().Add(duration)

	// Вывод точного времени
	fmt.Println("Точное время:", exactTime)
}
