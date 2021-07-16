package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var strForTask6 = "\ntask6 is a tool to check if tickets is lucky.\n" +
	"Flags:\n" +
	"\tflag -h displays help (-h);\n" +
	"\tflag -f path to file (-f=C:/elemtasks/task4/test.txt);\n" +
	"\tflag -m mode or 'Moscow' or 'Piter' (-m=Moscow)."

var fFile, fMode string
var fHelp bool

const noData = "noData"

func main() {
	//fillTickets()

	/*Флажки.*/
	flag.BoolVar(&fHelp, "h", false, "help info")
	flag.StringVar(&fFile, "f", noData, "path to file")
	flag.StringVar(&fMode, "m", noData, "mode or 'Moscow' or 'Piter'")
	flag.Parse()

	if fHelp || fFile == noData || fMode == noData || (fMode != "Moscow" && fMode != "Piter") {
		fmt.Println(strForTask6)
		return
	}

	tickets, err := getTickets()
	if err != nil {
		fmt.Println(err)
		return
	}

	luckyTickets, err := checkLuckyTickets(tickets)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(prepareAnswer(luckyTickets))
}

func getTickets() ([]string, error) {
	bytes, err := ioutil.ReadFile(fFile)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}

func prepareAnswer(luckyTickets []string) (answer string) {
	return fmt.Sprintf("[%s mode]. Lucky tickets in file = [%d]\n"+
		"Lucky array > %v", fMode, len(luckyTickets), luckyTickets)
}

func checkLuckyTickets(tickets []string) ([]string, error) {
	luckyTickets := make([]string, 0, 8)
	for _, strTicket := range tickets {
		forFirstNumber := map[bool]string{
			true:  strTicket[:3],
			false: string(strTicket[0]) + string(strTicket[2]) + string(strTicket[4]),
		}
		forSecondNumber := map[bool]string{
			true:  strTicket[3:],
			false: string(strTicket[1]) + string(strTicket[3]) + string(strTicket[5]),
		}

		n1, err := getSum(forFirstNumber[fMode == "Moscow"])
		if err != nil {
			return nil, err
		}
		n2, err := getSum(forSecondNumber[fMode == "Moscow"])
		if err != nil {
			return nil, err
		}

		if n1 == n2 {
			luckyTickets = append(luckyTickets, strTicket)
		}
	}

	return luckyTickets, nil
}

func getSum(strNumber string) (int, error) {
	var sum int
	for _, n := range strNumber {
		i, err := strconv.Atoi(string(n))
		if err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}
