package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main1() {
	//fmt.Println(task0(4, 8, "^"))
	//fmt.Printf("%v\n", task1())
	task2()
}

func task0(i1, i2 int, ch string) (chessboard string) {
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
	return
}

func task1() (arr []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for _, n := range strings.Split(scanner.Text(), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Printf("Wrong number: %s!\n", n)
			continue
		}
		if i > 0 && i%2 == 0 {
			arr = append(arr, i)
		}
	}
	return
}

func task2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cardNumber := strings.ReplaceAll(scanner.Text(), " ", "")
	/* https://golang.org/pkg/regexp/syntax/
	x => single character
	[[:digit:]] => digits (== [0-9])
	x{n} => exactly n x */
	if match, _ := regexp.MatchString(`^4[[:digit:]]{15}$`, cardNumber); match {
		fmt.Printf("**** **** **** %v\n", cardNumber[12:])
	} else {
		log.Println("Wrong card number!")
	}
}
