package main

/*
*Syntax Go. Homework 3
*Raisa Nenasheva, dated Apr 29, 2019
 */

import "fmt"

type car struct {
	model   string
	year    int
	valume  float64
	motor   bool
	windows bool
	filling float64
}
type truck struct {
	model   string
	year    int
	valume  float64
	motor   bool
	windows bool
	filling float64
}

func main() {
	var firstCar = car{model: "ford", year: 2012, valume: 421, motor: true,
		windows: false, filling: 55.5}
	var secondCar = car{model: "opel", year: 2015, valume: 400, motor: true,
		windows: false, filling: 100}
	var firstTruck = truck{model: "uaz", year: 2000, valume: 9000, motor: false,
		windows: true, filling: 70}
	var secondTruck = truck{model: "volvo", year: 2017, valume: 10000, motor: true,
		windows: true, filling: 90}
	fmt.Println(firstCar)
	fmt.Println(secondCar)
	fmt.Println(firstTruck)
	fmt.Println(secondTruck)
}
