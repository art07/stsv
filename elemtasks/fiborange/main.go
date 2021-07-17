package main

import (
	"flag"
	"fmt"
	"strings"
)

var strForTask8 = "\ntask8 is a tool which allows to display all Fibonacci numbers that are in the specified range.\n" +
	"arg1 <= arg2. Min value => 1, max value => 7778742049.\n" +
	"Flags:\n" +
	"\tflag > -h help info (-h);\n" +
	"\tflag > -n1 number1 (-n1=15);\n" +
	"\tflag > -n2 number2 (-n2=250)."

var number1, number2 uint64
var help bool

const minNum = 1
const maxNum = 7778742049

func main() {
	/*Флажки.*/
	flag.BoolVar(&help, "h", false, "help info")
	flag.Uint64Var(&number1, "n1", 0, "first number")
	flag.Uint64Var(&number2, "n2", 0, "second number")
	flag.Parse()

	/*Проверки.*/
	if help || number1 < minNum || number2 < minNum || number1 > number2 || number1 > maxNum || number2 > maxNum {
		fmt.Println(strForTask8)
		return
	}

	/*Последовательность фибоначи до 7778742049*/
	nextFiboValue := fibonacciFunc()
	var fiboArray [50]uint64
	for i := 0; i < 50; i++ {
		fiboArray[i] = nextFiboValue()
	}

	/*Числа фибоначи определенного диапазона.*/
	var strBuilder strings.Builder
	_, _ = fmt.Fprintf(&strBuilder, "Range [%d] - [%d] > \n", number1, number2)
	for _, n := range fiboArray {
		if n >= number1 && n <= number2 {
			_, _ = fmt.Fprintf(&strBuilder, "%d,", n)
		}
	}

	/*Вывод результата.*/
	fmt.Println(strings.TrimRight(strBuilder.String(), ","))
}

func fibonacciFunc() func() uint64 {
	var num1, num2, nextFibo, loop uint64
	num2 = 1
	return func() uint64 {
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
