package main

import (
	"errors"
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

	if err := setFlags(&fHelp, &fFile, &fSearch, &fNewWord); err != nil {
		fmt.Println(err)
		fmt.Println(strForTask4)
		return
	}

	if fNewWord == noData {
		i, err := numberOfOccurrences(fFile, fSearch)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Mode #1\n[%s] meets %d time(s)\n", fSearch, i)
		return
	}
	fmt.Println("Mode #2")
	if err := changeText(fFile, fSearch, fNewWord); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mode #2 finished!")
}

func setFlags(fHelp *bool, fFile, fSearch, fNewWord *string) error {
	flag.BoolVar(fHelp, "h", false, "help info")
	flag.StringVar(fFile, "f", noData, "path to file")
	flag.StringVar(fSearch, "s", noData, "search word")
	flag.StringVar(fNewWord, "n", noData, "new word")
	flag.Parse()

	if *fHelp || *fFile == noData || *fSearch == noData {
		return errors.New("not enough flags data to proceed or need help")
	}
	return nil
}

func numberOfOccurrences(filePath, searchWord string) (int, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	return strings.Count(string(buf), searchWord), nil
}

func changeText(fFile, fSearch, fNewWord string) error {
	buf, err := ioutil.ReadFile(fFile)
	if err != nil {
		return err
	}
	str := strings.ReplaceAll(string(buf), fSearch, fNewWord)
	if err = os.WriteFile(fFile, []byte(str), 0); err != nil {
		return err
	}
	return nil
}
