package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var strForTask4 = "\ntask4 is a tool for jobs with files and strings.\n" +
	"Works in two modes. Mode #1 > count number of occurrences; Mode #2 > text replacement.\n" +
	"If -c == noData > Mode #1, else Mode #2.\n" +
	"Flags:\n" +
	"\tflag -h displays help (-h);\n" +
	"\tflag -f path to file (-f=C:/elemtasks/task4/test.txt);\n" +
	"\tflag -s search word (-s=Lorem);\n" +
	"\tflag -n new word (-n=Lorem123)."

var fFile, fSearch, fNewWord string
var fHelp bool

const noData = "noData"

func main() {
	/*Флаги.*/
	flag.BoolVar(&fHelp, "h", false, "help info")
	flag.StringVar(&fFile, "f", noData, "path to file")
	flag.StringVar(&fSearch, "s", noData, "search word")
	flag.StringVar(&fNewWord, "n", noData, "new word")
	flag.Parse()

	if fHelp || fFile == noData || fSearch == noData {
		fmt.Println(strForTask4)
		return
	}

	if fNewWord == noData {
		i, err := numberOfOccurrences()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Mode #1\n[%s] meets %d time(s)\n", fSearch, i)
	} else {
		fmt.Println("Mode #2")
		if err := changeText(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Mode #2 finished!")
	}
}

func numberOfOccurrences() (int, error) {
	buf, err := ioutil.ReadFile(fFile)
	if err != nil {
		return 0, err
	}
	return strings.Count(string(buf), fSearch), nil
}

func changeText() error {
	buf, err := ioutil.ReadFile(fFile)
	if err != nil {
		return err
	}
	str := strings.ReplaceAll(string(buf), fSearch, fNewWord)
	if err = ioutil.WriteFile(fFile, []byte(str), 0); err != nil {
		return err
	}
	return nil
}
