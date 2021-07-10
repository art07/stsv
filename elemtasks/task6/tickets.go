package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var strForTask6 = "\ntask6 is a tool to check if ticket is lucky.\n" +
	"Args:\n" +
	"\targ1> path to file with tickets (C:/elemtasks/task6);\n" +
	"\targ2> mode ('Moscow' or 'Piter')."

func main() {
	//fillTickets()
	if len(os.Args) != 3 || (os.Args[2] != "Moscow" && os.Args[2] != "Piter") {
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
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}

func prepareAnswer(luckyTickets []string) (answer string) {
	return fmt.Sprintf("[%s mode]. Lucky tickets in file = [%d]\n"+
		"Lucky array > %v", os.Args[2], len(luckyTickets), luckyTickets)
}

func checkLuckyTickets(tickets []string) ([]string, error) {
	var luckyTickets []string

	for _, strTicket := range tickets {
		forFirstNumber := map[bool]string{
			true:  strTicket[:3],
			false: string(strTicket[0]) + string(strTicket[2]) + string(strTicket[4]),
		}
		forSecondNumber := map[bool]string{
			true:  strTicket[3:],
			false: string(strTicket[1]) + string(strTicket[3]) + string(strTicket[5]),
		}

		n1, err := getSum(forFirstNumber[os.Args[2] == "Moscow"])
		if err != nil {
			return nil, err
		}
		n2, err := getSum(forSecondNumber[os.Args[2] == "Moscow"])
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
