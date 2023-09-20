package main

import "fmt"

// Определяем интерфейс обработчика запроса
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleRequest(request string) {
	// Если есть следующий обработчик, передаем ему запрос
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

// Реализация конкретного обработчика запроса
type ConcreteHandler1 struct {
	BaseHandler
}

func (c *ConcreteHandler1) HandleRequest(request string) {
	// Проверяем, может ли этот обработчик обработать запрос
	if request == "handler1" {
		fmt.Println("Обработчик 1 обрабатывает запрос")
	} else {
		// Если не может, передаем запрос следующему обработчику
		fmt.Println("Обработчик 1 передает запрос следующему обработчику")
		c.BaseHandler.HandleRequest(request)
	}
}

// Реализация другого обработчика запроса
type ConcreteHandler2 struct {
	BaseHandler
}

func (c *ConcreteHandler2) HandleRequest(request string) {
	if request == "handler2" {
		fmt.Println("Обработчик 2 обрабатывает запрос")
	} else {
		fmt.Println("Обработчик 2 передает запрос следующему обработчику")
		c.BaseHandler.HandleRequest(request)
	}
}

func main() {
	// Создаем цепочку обработчиков
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler1.SetNext(handler2)

	// Обрабатываем запросы
	handler1.HandleRequest("handler1")
	handler1.HandleRequest("handler2")
	handler1.HandleRequest("handler3")
}
