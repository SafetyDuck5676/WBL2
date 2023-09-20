package main

import "fmt"

// Унифицированный интерфейс, предоставляемый фасадом
type ShapeFacade interface {
	Draw()
}

// Внешний компонент 1
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Рисую окружность")
}

// Внешний компонент 2
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Рисую квадрат")
}

// Внешний компонент 3
type Triangle struct{}

func (t *Triangle) Draw() {
	fmt.Println("Рисую треугольник")
}

// Фасад, предоставляющий упрощенный интерфейс для клиента
type DrawingFacade struct {
	Circle   ShapeFacade
	Square   ShapeFacade
	Triangle ShapeFacade
}

func (d *DrawingFacade) DrawCircle() {
	d.Circle.Draw()
}

func (d *DrawingFacade) DrawSquare() {
	d.Square.Draw()
}

func (d *DrawingFacade) DrawTriangle() {
	d.Triangle.Draw()
}

// Пример использования
func main() {
	circle := &Circle{}
	square := &Square{}
	triangle := &Triangle{}
	facade := &DrawingFacade{
		Circle:   circle,
		Square:   square,
		Triangle: triangle,
	}

	// Используем фасад для рисования
	facade.DrawCircle()
	facade.DrawSquare()
	facade.DrawTriangle()
}
