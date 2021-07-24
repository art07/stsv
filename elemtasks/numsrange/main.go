package main

import (
	"errors"
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
	parseFlags(&fHelp, &fNumber)

	if err := flagValidation(fHelp, fNumber); err != nil {
		fmt.Println(err, strForTask7)
		return
	}

	fmt.Println(getRange(fNumber))
}

func parseFlags(fHelp *bool, fNumber *int) {
	flag.BoolVar(fHelp, "h", false, "help info")
	flag.IntVar(fNumber, "n", 0, "input number")
	flag.Parse()
}

func flagValidation(fHelp bool, fNumber int) error {
	if fHelp || fNumber < 1 {
		return errors.New("interrupt the app because of flags")
	}
	return nil
}

func getRange(fNumber int) string {
	var strBuild strings.Builder
	for i := 1; ; i++ {
		if (i * i) < fNumber {
			strBuild.WriteString(fmt.Sprintf("%d,", i))
			continue
		}
		break
	}
	return strings.TrimRight(strBuild.String(), ",")
}
