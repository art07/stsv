package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var strForTask8 = "\ntask8 is a tool which allows to display all Fibonacci numbers that are in the specified range.\n" +
	"arg1 <= arg2. Min value => 1, max value => 7778742049.\n" +
	"Args:\n" +
	"\targ1> number (1-7778742049);\n" +
	"\targ2> number (1-7778742049)."

func main() {
	/*Проверки.*/
	if len(os.Args) < 3 {
		fmt.Println(strForTask8)
		return
	}
	i1, i2, err := getUints()
	if err != nil {
		fmt.Println(err)
		return
	}
	if i1 > i2 || i1 < 1 || i2 < 1 || i1 > 7778742049 || i2 > 7778742049 {
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
	str := fmt.Sprintf("Range [%d] - [%d] > \n", i1, i2)
	for _, number := range fiboArray {
		if number >= i1 && number <= i2 {
			str += fmt.Sprintf("%d,", number)
		}
	}

	/*Вывод результата.*/
	fmt.Println(strings.TrimRight(str, ","))
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

func getUints() (i1 uint64, i2 uint64, err error) {
	// https://golang.org/pkg/strconv/#ParseUint
	i1, err = strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		return
	}
	i2, err = strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		return
	}
	return
}
