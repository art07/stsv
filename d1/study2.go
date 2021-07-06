package main

import "fmt"

// https://ru.wikipedia.org/wiki/Числа_Фибоначчи
func fibonacciFunc() func() int {
	var num1, nextFibo, loop int
	num2 := 1
	return func() int {
		switch loop {
		case 0:
		case 1:
			nextFibo = 1
		default:
			nextFibo = num1 + num2
			num1 = num2
			num2 = nextFibo
		}
		loop++
		return nextFibo
	}
}

func main() {
	nextFiboValue := fibonacciFunc()
	for i := 0; i < 40; i++ {
		fmt.Printf("%d ", nextFiboValue())
	}
}
