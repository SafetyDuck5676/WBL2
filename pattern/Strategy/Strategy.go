package main

import "fmt"

// Абстрактный класс "DiscountStrategy" с методом "CalculateDiscount"
type DiscountStrategy interface {
	CalculateDiscount(order *Order) float64
}

// Конкретные классы, реализующие интерфейс "DiscountStrategy"
type BaseDiscountStrategy struct{}

func (s *BaseDiscountStrategy) CalculateDiscount(order *Order) float64 {
	return order.TotalAmount() * 0.1
}

type LoyaltyDiscountStrategy struct{}

func (s *LoyaltyDiscountStrategy) CalculateDiscount(order *Order) float64 {
	return order.TotalAmount() * 0.2
}

type PromoCodeDiscountStrategy struct{}

func (s *PromoCodeDiscountStrategy) CalculateDiscount(order *Order) float64 {
	return order.TotalAmount() * 0.15
}

// Класс "Order" с полем для хранения выбранной стратегии расчета скидки и методами для применения скидки
type Order struct {
	discountStrategy DiscountStrategy
}

func (o *Order) SetDiscountStrategy(strategy DiscountStrategy) {
	o.discountStrategy = strategy
}

func (o *Order) CalculateTotalAmount() float64 {
	totalAmount := o.TotalAmount()

	if o.discountStrategy != nil {
		discount := o.discountStrategy.CalculateDiscount(o)
		totalAmount -= discount
	}

	return totalAmount
}

func (o *Order) TotalAmount() float64 {
	return 100.0 // Пример значения, обычно считается более сложно
}

func main() {
	order := &Order{}
	order.SetDiscountStrategy(&BaseDiscountStrategy{})

	totalAmount := order.CalculateTotalAmount()
	fmt.Printf("Общая сумма заказа с примененной скидкой: %.2f\n", totalAmount)

	order.SetDiscountStrategy(&LoyaltyDiscountStrategy{})

	totalAmount = order.CalculateTotalAmount()
	fmt.Printf("Общая сумма заказа с примененной скидкой: %.2f\n", totalAmount)
}
