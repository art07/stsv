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
	"Works in two modes. Mode #1 > count number of occurrences; " +
	"Mode #2 > text replacement.\n" +
	"If -c == noData > Mode #1, else Mode #2.\n" +
	"Flags:\n" +
	"\tflag -h displays help (-h);\n" +
	"\tflag -f path to file (-f=C:/elemtasks/task4/text.txt);\n" +
	"\tflag -s search word (-s=Lorem);\n" +
	"\tflag -n new word (-n=Lorem123)."

const noData = "noData"
const modeStr = "Mode #%d finished successfully! [%s] meets = %d time(s)."

func main() {
	var fFile, fSearch, fNewWord string
	var fHelp bool
	parseFlags(&fHelp, &fFile, &fSearch, &fNewWord)

	if err := flagValidation(fHelp, fFile, fSearch); err != nil {
		fmt.Println(err, strForTask4)
		return
	}

	switch fNewWord {
	case noData:
		fmt.Println(mode1(fFile, fSearch))
	default:
		fmt.Println(mode2(fFile, fSearch, fNewWord))
	}
}

func parseFlags(fHelp *bool, fFile, fSearch, fNewWord *string) {
	flag.BoolVar(fHelp, "h", false, "help info")
	flag.StringVar(fFile, "f", noData, "path to file")
	flag.StringVar(fSearch, "s", noData, "search word")
	flag.StringVar(fNewWord, "n", noData, "new word")
	flag.Parse()
}

func flagValidation(fHelp bool, fFile, fSearch string) error {
	if fHelp || fFile == noData || fSearch == noData {
		return errors.New("interrupt the app because of flags")
	}
	return nil
}

func mode1(fFile string, fSearch string) string {
	fmt.Println("Mode #1")
	if i, err := numberOfOccurrences(fFile, fSearch); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf(modeStr, 1, fSearch, i)
	}
}

func mode2(fFile string, fSearch string, fNewWord string) string {
	fmt.Println("Mode #2")
	if i, err := changeText(fFile, fSearch, fNewWord); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf(modeStr, 2, fNewWord, i)
	}
}

func numberOfOccurrences(filePath, searchWord string) (int, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	return strings.Count(string(buf), searchWord), nil
}

func changeText(fFile, fSearch, fNewWord string) (int, error) {
	buf, err := ioutil.ReadFile(fFile)
	if err != nil {
		return 0, err
	}

	str := strings.ReplaceAll(string(buf), fSearch, fNewWord)
	_ = os.WriteFile(fFile, []byte(str), 0)

	i, _ := numberOfOccurrences(fFile, fNewWord)
	return i, nil
}
