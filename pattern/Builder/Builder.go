package main

import "fmt"

// Структура для представления сложного объекта - автомобиля
type Car struct {
	Brand  string
	Model  string
	Engine string
	Wheels int
}

// Интерфейс, определяющий шаги построения автомобиля
type CarBuilder interface {
	SetBrand(brand string)
	SetModel(model string)
	SetEngine(engine string)
	SetWheels(wheels int)
	GetCar() Car
}

// Конкретная реализация строителя для автомобиля
type ConcreteCarBuilder struct {
	car Car
}

func (b *ConcreteCarBuilder) SetBrand(brand string) {
	b.car.Brand = brand
}

func (b *ConcreteCarBuilder) SetModel(model string) {
	b.car.Model = model
}

func (b *ConcreteCarBuilder) SetEngine(engine string) {
	b.car.Engine = engine
}

func (b *ConcreteCarBuilder) SetWheels(wheels int) {
	b.car.Wheels = wheels
}

func (b *ConcreteCarBuilder) GetCar() Car {
	return b.car
}

// Директор, который управляет процессом построения автомобиля
type Director struct {
	builder CarBuilder
}

func (d *Director) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

func (d *Director) Construct() Car {
	d.builder.SetBrand("Tesla")
	d.builder.SetModel("Model S")
	d.builder.SetEngine("Electric")
	d.builder.SetWheels(4)
	return d.builder.GetCar()
}

func main() {
	var builder CarBuilder = &ConcreteCarBuilder{} // Создаем конкретного строителя
	var director Director = Director{}             // Создаем директора
	director.SetBuilder(builder)                   // Назначаем строителя директору
	car := director.Construct()                    // Строим автомобиль с помощью директора и строителя

	fmt.Printf("Brand: %s\nModel: %s\nEngine: %s\nWheels: %d\n",
		car.Brand, car.Model, car.Engine, car.Wheels)
}
