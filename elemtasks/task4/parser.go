package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var args = os.Args[1:]

var strForTask4 = "\ntask4 is a tool for jobs with files.\n" +
	"Works in two modes. With two arguments, the number of occurrences; with three, line replacement.\n" +
	"Mode1. Args:\n" +
	"\targ1 pathToFile (С:/docs/test.txt);\n" +
	"\targ2 search string (\"any string\").\n" +
	"Mode2. Args:\n" +
	"\targ1 pathToFile (С:/docs/test.txt);\n" +
	"\targ2 search string (\"any string\");\n" +
	"\targ3 change string (\"any string\")."

//goland:noinspection GoPrintFunctions
func main() {
	switch len(args) {
	case 2:
		i, err := numberOfOccurrences()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Mode #1")
		fmt.Printf("[%s] meets %d time(s)\n", args[1], i)
	case 3:
		fmt.Println("Mode #2")
		err := changeText()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Mode #2 finished!")
	default:
		fmt.Println(strForTask4)
	}
}

//goland:noinspection GoUnhandledErrorResult
func numberOfOccurrences() (int, error) {
	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		return 0, err
	}
	return strings.Count(string(buf), args[1]), nil
}

//goland:noinspection GoUnhandledErrorResult
func changeText() error {
	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}
	str := strings.ReplaceAll(string(buf), args[1], args[2])
	if err = ioutil.WriteFile(args[0], []byte(str), 0); err != nil {
		return err
	}
	return nil
}
