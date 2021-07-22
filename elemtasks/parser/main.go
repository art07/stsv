package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var strForTask4 = "\ntask4 is a tool for jobs with files and strings.\n" +
	"Works in two modes. Mode #1 > count number of occurrences; Mode #2 > text replacement.\n" +
	"If -c == noData > Mode #1, else Mode #2.\n" +
	"Flags:\n" +
	"\tflag -h displays help (-h);\n" +
	"\tflag -f path to file (-f=C:/elemtasks/task4/text.txt);\n" +
	"\tflag -s search word (-s=Lorem);\n" +
	"\tflag -n new word (-n=Lorem123)."

const noData = "noData"

func main() {
	var fFile, fSearch, fNewWord string
	var fHelp bool
	flag.BoolVar(&fHelp, "h", false, "help info")
	flag.StringVar(&fFile, "f", noData, "path to file")
	flag.StringVar(&fSearch, "s", noData, "search word")
	flag.StringVar(&fNewWord, "n", noData, "new word")
	flag.Parse()

	if ok := interrupt(fHelp, fFile, fSearch); ok {
		return
	}

	if fNewWord == noData {
		i, err := numberOfOccurrences(fFile, fSearch)
		if err != nil {
			return
		}
		fmt.Printf("Mode #1\n[%s] meets %d time(s)\n", fSearch, i)
		return
	}

	if err := changeText(fFile, fSearch, fNewWord); err != nil {
		return
	}
}

func interrupt(fHelp bool, fFile, fSearch string) bool {
	if fHelp || fFile == noData || fSearch == noData {
		fmt.Println(strForTask4)
		return true
	}
	return false
}

func numberOfOccurrences(filePath, searchWord string) (int, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return strings.Count(string(buf), searchWord), nil
}

func changeText(fFile, fSearch, fNewWord string) error {
	fmt.Println("Mode #2")
	buf, err := ioutil.ReadFile(fFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	str := strings.ReplaceAll(string(buf), fSearch, fNewWord)
	//goland:noinspection GoUnhandledErrorResult
	os.WriteFile(fFile, []byte(str), 0)
	fmt.Println("Mode #2 finished successfully!")
	return nil
}
