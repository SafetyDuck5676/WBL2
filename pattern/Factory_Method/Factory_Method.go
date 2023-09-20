package main

import "fmt"

// Интерфейс, который определяет метод создания продукта
type Creator interface {
	CreateProduct() Product
}

// Интерфейс, который определяет метод использования продукта
type Product interface {
	Use()
}

// Реализация продукта A
type ProductA struct{}

func (p ProductA) Use() {
	fmt.Println("Using Product A")
}

// Реализация продукта B
type ProductB struct{}

func (p ProductB) Use() {
	fmt.Println("Using Product B")
}

// Фабричный метод для продукта A
type ConcreteCreatorA struct{}

func (c ConcreteCreatorA) CreateProduct() Product {
	return ProductA{}
}

// Фабричный метод для продукта B
type ConcreteCreatorB struct{}

func (c ConcreteCreatorB) CreateProduct() Product {
	return ProductB{}
}

func main() {
	// Создание объекта через фабричный метод ConcreteCreatorA
	creatorA := ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	productA.Use()

	// Создание объекта через фабричный метод ConcreteCreatorB
	creatorB := ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	productB.Use()
}
