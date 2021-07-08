package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var args = os.Args[1:]

var strForTask4 = "\ntask4 is a tool for jobs with files.\n" +
	"Works in two modes. With two arguments, the number of occurrences; with three, line replacement.\n" +
	"Mode1. Args:\n" +
	"\targ1 pathToFile (С:/docs/test.txt);\n" +
	"\targ2 search string (any string).\n" +
	"Mode2. Args:\n" +
	"\targ1 pathToFile (С:/docs/test.txt);\n" +
	"\targ2 search string (any string);\n" +
	"\targ3 change string (any string)."

//goland:noinspection GoPrintFunctions
func main() {
	switch len(args) {
	case 2:
		fmt.Println("Mode #1")
		i, err := numberOfOccurrences()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("[%s] meets %d time(s)\n", args[1], i)
	case 3:
		fmt.Println("Mode #2")
		// do some job ...
	default:
		fmt.Println(strForTask4)
	}
}

//goland:noinspection GoUnhandledErrorResult
func numberOfOccurrences() (int, error) {
	file, err := os.Open(args[0])
	if err != nil {
		return 0, err
	}
	defer file.Close()

	snr := bufio.NewScanner(file)
	times := 0
	for snr.Scan() {
		times += strings.Count(snr.Text(), args[1])
	}
	if snr.Err() != nil {
		return times, snr.Err()
	}
	return times, nil
}
