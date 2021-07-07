package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var strForTask1 = "\ntask1 is a tool for printing chessboard.\n" +
	"Args:\n" +
	"\targ1 is number of rows (4);\n" +
	"\targ2 is number of columns (8);\n" +
	"\targ3 one of [#$%&*+@_~^] (\"*\")."

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println(strForTask1)
		return
	}

	if len(os.Args) != 4 {
		log.Println("os.Args len != 4")
		return
	}
	i1, i2, ch, err := getData(os.Args[1:])
	if err != nil {
		log.Println("Get data from os.Args failed")
		return
	}

	chessboard := ""
	for i := 1; i <= i1; i++ {
		for j := 0; j < i2; j++ {
			chessboard += fmt.Sprintf("%s ", ch)
		}
		if (i % 2) != 0 {
			chessboard += "\n "
		} else {
			chessboard += "\n"
		}
	}
	fmt.Println(chessboard)
}

func getData(args []string) (int, int, string, error) {
	i1, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, "", err
	}
	i2, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, "", err
	}
	match, err := regexp.MatchString(`^[#$%&*+@_~^]$`, args[2])
	if !match || err != nil {
		return 0, 0, "", errors.New(fmt.Sprintf("args[2] error; <%s> not fit", args[2]))
	}

	return i1, i2, args[2], nil
}
