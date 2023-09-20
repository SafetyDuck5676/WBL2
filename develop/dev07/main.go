package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция `or` объединяет один или более done-каналов в single-канал.
// Если один из done-каналов закрывается, функция возвращает single-канал.
// Входные параметры функции - один или несколько done-каналов.
// Возвращаемое значение - single-канал.
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаем канал, через который будем передавать сигналы
	// о закрытии done-каналов
	orDone := make(chan interface{})

	// Создаем горутину, которая будет слушать состояние done-каналов
	go func() {
		// Отложенная функция для закрытия канала orDone
		defer close(orDone)

		// Создаем канал для ожидания готовности работы с done-каналами
		var waiting sync.WaitGroup

		// Функция, которая слушает состояние одного done-канала
		// и передает сигнал о закрытии в канал orDone
		recv := func(ch <-chan interface{}) {
			defer waiting.Done()

			// Ожидаем, пока done-канал не будет закрыт
			<-ch
		}

		// Запускаем горутину для каждого done-канала
		for _, ch := range channels {
			waiting.Add(1)
			go recv(ch)
		}

		// Ожидаем, пока все горутины закончат свою работу
		waiting.Wait()
	}()

	// Возвращаем single-канал
	return orDone
}

func main() {
	// Функция, которая создает done-канал с заданным временем задержки
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	// Используем функцию or для объединения нескольких done-каналов
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
