package main

import "fmt"

// Интерфейс посетителя, определяющий различные операции над элементами
type Visitor interface {
	VisitElementA(elementA *ElementA)
	VisitElementB(elementB *ElementB)
}

// Конкретный посетитель 1
type Visitor1 struct{}

func (v *Visitor1) VisitElementA(elementA *ElementA) {
	fmt.Println("Visitor1: VisitElementA")
}

func (v *Visitor1) VisitElementB(elementB *ElementB) {
	fmt.Println("Visitor1: VisitElementB")
}

// Конкретный посетитель 2
type Visitor2 struct{}

func (v *Visitor2) VisitElementA(elementA *ElementA) {
	fmt.Println("Visitor2: VisitElementA")
}

func (v *Visitor2) VisitElementB(elementB *ElementB) {
	fmt.Println("Visitor2: VisitElementB")
}

// Интерфейс элемента, определяющий метод accept, принимающий посетителя
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ElementA struct{}

func (a *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(a)
}

// Конкретный элемент B
type ElementB struct{}

func (b *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(b)
}

// Клиентский код
func main() {
	elements := []Element{
		&ElementA{},
		&ElementB{},
	}

	visitor1 := &Visitor1{}
	visitor2 := &Visitor2{}

	for _, element := range elements {
		element.Accept(visitor1)
		element.Accept(visitor2)
	}
}
