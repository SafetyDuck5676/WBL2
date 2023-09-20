package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Создаем done-каналы
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	// Функция, которая закрывает done-канал после заданной задержки
	sig := func(ch chan interface{}, after time.Duration) {
		go func() {
			time.Sleep(after)
			close(ch)
		}()
	}

	// Запускаем тест на функцию or
	start := time.Now()
	go func() {
		<-or(ch1, ch2, ch3)
		t.Logf("Done after %v\n", time.Since(start))
	}()

	// Закрываем done-каналы поочереди
	sig(ch1, 1*time.Second)
	sig(ch2, 2*time.Second)
	sig(ch3, 3*time.Second)

	// Ждем завершения теста
	time.Sleep(4 * time.Second)
}
