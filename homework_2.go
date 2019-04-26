package main

/*
*Syntax Go. Homework 2
*Raisa Nenasheva, dated Apr 26, 2019
 */

import (
	"fmt"
)

func main() {
	var num int
	var n int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	evenNumber(num)
	divideByThree(num)
	fmt.Println("Введите длину последовательности чисел Фибоначчи")
	fmt.Scanln(&n)
	fiboNum := fibonacciNumbers()
	for i := 0; i < n; i++ {
		switch {
		case i == 0:
			fmt.Println(i)
		default:
			fmt.Println(fiboNum())
		}
	}
}

func evenNumber(num int) {
	if num%2 == 0 {
		fmt.Println("Вы ввели четное число")
	} else {
		fmt.Println("Вы ввели нечетное число")
	}
}
func divideByThree(num int) {
	if num%3 == 0 {
		fmt.Println("Данное число делется на три без остатка")
	} else {
		fmt.Println("Данное число нельзя поделить на три без остатка")
	}
}
func fibonacciNumbers() func() int {
	fibo1 := 0
	fibo2 := 1
	return func() int {
		fibo := fibo1 + fibo2
		fibo1 = fibo2
		fibo2 = fibo
		return fibo1
	}
}
