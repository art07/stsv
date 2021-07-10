package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var strForTask7 = "\ntask7 is a tool to find numbers that in square less than given number.\n" +
	"Min value => 1.\n" +
	"Args:\n" +
	"\targ1> number (150)."

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strForTask7)
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number < 1 {
		fmt.Println(strForTask7)
		return
	}
	var str string
	for i := 1; ; i++ {
		if (i * i) < number {
			str += fmt.Sprintf("%d,", i)
		} else {
			break
		}
	}
	fmt.Println(strings.TrimRight(str, ","))
}
