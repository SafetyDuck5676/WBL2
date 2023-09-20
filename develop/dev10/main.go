package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Функция для обработки ошибок
func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	// Получение хоста и порта из аргументов командной строки
	host := flag.Arg(0)
	port := flag.Arg(1)

	// Проверка, что указаны хост и порт
	if host == "" || port == "" {
		fmt.Fprintln(os.Stderr, "Host and port must be specified")
		os.Exit(1)
	}

	// Установка соединения с указанным хостом и портом
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	handleError(err)
	defer conn.Close()

	// Установка обработчика сигнала для завершения программы при нажатии Ctrl+C
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Запуск горутины для чтения данных из сокета и вывода их в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil && !errors.Is(err, io.EOF) {
			handleError(err)
		}
		signalCh <- os.Interrupt // Отправка сигнала завершения программы
	}()

	// Копирование данных из STDIN в сокет
	_, err = io.Copy(conn, os.Stdin)
	handleError(err)
}
