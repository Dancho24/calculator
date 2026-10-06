package main

import "fmt"

const (
	BaseRate     float64 = 5.50
	TaxRate      float64 = 0.12
	DistanceRate float64 = 2.0
	FragileFee   float64 = 0.2
)

func main() {
	var name string
	var weight float64
	var distance float64
	var fragileCount float64

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите вес груза (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Введите дистанцию (км): ")
	fmt.Scan(&distance)

	fmt.Print("Введите количество хрупких упаковок: ")
	fmt.Scan(&fragileCount)

	baseCost := (weight*BaseRate)*(1+FragileFee*fragileCount) + (distance * DistanceRate)

	totalCost := baseCost * (1 + TaxRate)

	fmt.Println("\n--- Итоговый отчет ---")
	fmt.Printf("Заказчик: %s\n", name)
	fmt.Printf("Базовая стоимость: %.2f\n", baseCost)
	fmt.Printf("Итоговая сумма (с налогом 12%%): %.2f\n", totalCost)
}

//
