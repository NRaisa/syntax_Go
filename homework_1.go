package main

import (
	"fmt"
	"math"
)

func main() {
	converter()
	triangle()
	bank()
}
func converter() {
	const dollar = 64.05
	var rub float64
	fmt.Println("Введите конвертируемую сумму в рублях:")
	fmt.Scanln(&rub)
	var convertesDollar = rub / dollar
	fmt.Println("Ваша сумма в долларах составит:", convertesDollar)
}
func triangle() {
	var a float64
	var b float64
	var c float64
	var s float64
	var p float64
	fmt.Println("Введите длину катета a (см)")
	fmt.Scanln(&a)
	fmt.Println("Введите длину катета b (см)")
	fmt.Scanln(&b)
	c = math.Pow(a, 2) + math.Pow(b, 2)
	c = math.Sqrt(c)
	s = 0.5 * a * b
	p = a + b + c
	fmt.Printf("Гипотенуза прямоугольного треугольника равна: %f см, площадь: %f см.кв., периметр: %f см \n", c, s, p)
}
func bank() {
	var sum float64
	var percent float64
	var sum1 float64
	var sum2 float64
	const k float64 = 100
	var year float64 = 365
	time := year * 5
	fmt.Println("Введите сумму вклада")
	fmt.Scanln(&sum)
	fmt.Println("Введите годовой процент по вкладу")
	fmt.Scanln(&percent)
	sum1 = sum + (sum*percent*time)/(time*k)
	sum2 = 1 + (percent*year)/(k*time)
	sum2 = math.Pow(sum2, 5)
	sum2 = sum * sum2
	fmt.Printf("Через 5 лет сумма на Вашем счету составит %f без капитализации начисленных процентов \n %f при ежегодной капитализации процентов", sum1, sum2)
}
