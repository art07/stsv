package main

import (
	"flag"
	"fmt"
	"strings"
)

var strForTask7 = "\ntask7 is a tool to find numbers that in square less than given number.\n" +
	"Min value => 1.\n" +
	"Flag:\n" +
	"\tflag > -h help info (-h);\n" +
	"\tflag > -n any number (-n=150)."

var fHelp bool
var fNumber int

func main() {
	/*Флажки.*/
	flag.BoolVar(&fHelp, "h", false, "help info")
	flag.IntVar(&fNumber, "n", 0, "input number")
	flag.Parse()

	if fHelp || fNumber < 1 {
		fmt.Println(strForTask7)
		return
	}

	var str string
	for i := 1; ; i++ {
		if (i * i) < fNumber {
			str += fmt.Sprintf("%d,", i)
			continue
		}
		break
	}
	fmt.Println(strings.TrimRight(str, ","))
}
