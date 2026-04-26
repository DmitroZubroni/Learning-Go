package main

import (
	"fmt"
	"math"
)

func main() {
	// объявление переменных
	const IMTPower = 2 // константа
	// userHeight := 1.8
	// userWeight := 100.0

	var userHeight, userWeight float64

	fmt.Print("калькулятор индекса массы тела _\n")

	fmt.Print("укажите свой рост в метрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("укажите свой вес: ")
	fmt.Scan(&userWeight)
	// явная конвертанция
	//var IMT = float64(userWeight) / userHeight

	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)

	// альтернативные способы объявление переменных
	//	userHeight := 1.8
	//	var userWeight float64
	// объявление нескольких переменных
	//	var userHeight, userWeight float64 = 1.8, 100
	//	userHeight, userWeight := 1.8, 100.0
}
