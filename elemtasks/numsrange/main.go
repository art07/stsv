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

func main() {
	var fHelp bool
	var fNumber int
	flag.BoolVar(&fHelp, "h", false, "help info")
	flag.IntVar(&fNumber, "n", 0, "input number")
	flag.Parse()

	if ok := interrupt(fHelp, fNumber); ok {
		fmt.Println(strForTask7)
		return
	}

	fmt.Println(getRange(fNumber))
}

func interrupt(fHelp bool, fNumber int) bool {
	if fHelp || fNumber < 1 {
		return true
	}
	return false
}

func getRange(fNumber int) string {
	var strBuild strings.Builder
	for i := 1; ; i++ {
		if (i * i) < fNumber {
			//goland:noinspection GoUnhandledErrorResult
			fmt.Fprintf(&strBuild, "%d,", i)
			continue
		}
		break
	}
	return strings.TrimRight(strBuild.String(), ",")
}
