package main

//"art/edu/stsv/randtick" [FillTickets function]

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var strForTask6 = "\ntask6 is a tool to check if tickets is lucky.\n" +
	"Flags:\n" +
	"\tflag -h displays help (-h);\n" +
	"\tflag -f path to file (-f=C:/elemtasks/task4/test.txt);\n" +
	"\tflag -m mode or 'Moscow' or 'Piter' (-m=Moscow)."

const noData = "noData"

func main() {
	var fFile, fMode string
	var fHelp bool
	parseFlags(&fHelp, &fFile, &fMode)

	if err := flagValidation(fHelp, fFile, fMode); err != nil {
		fmt.Println(err, strForTask6)
		return
	}

	tickets, err := getTickets(fFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tickets) == 0 {
		fmt.Println("No tickets to find lucky")
		return
	}

	luckyTickets := checkLuckyTickets(tickets, fMode)
	if len(luckyTickets) == 0 {
		fmt.Println("No lucky tickets.")
		return
	}
	fmt.Printf("[%s mode]. Lucky tickets => [%d]\nLucky array > %v\n", fMode, len(luckyTickets), luckyTickets)
}

func parseFlags(fHelp *bool, fFile, fMode *string) {
	flag.BoolVar(fHelp, "h", false, "help info")
	flag.StringVar(fFile, "f", noData, "path to file")
	flag.StringVar(fMode, "m", noData, "mode or 'Moscow' or 'Piter'")
	flag.Parse()
}

func flagValidation(fHelp bool, fFile string, fMode string) error {
	if fHelp || fFile == noData || fMode == noData || (fMode != "Moscow" && fMode != "Piter") {
		return errors.New("interrupt the app because of flags")
	}
	return nil
}

func getTickets(filePath string) (tickets []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	snr := bufio.NewScanner(file)
	for snr.Scan() {
		ticket := strings.TrimRight(snr.Text(), "\n")
		if matched, _ := isTicket(ticket); matched {
			tickets = append(tickets, ticket)
		}
	}
	return
}

func isTicket(ticket string) (bool, error) {
	return regexp.MatchString(`^[[:digit:]]{6}$`, ticket)
}

func checkLuckyTickets(tickets []string, mode string) (luckyTickets []string) {
	for _, strTicket := range tickets {
		forFirstNumber := map[bool]string{
			true:  strTicket[:3],
			false: string(strTicket[0]) + string(strTicket[2]) + string(strTicket[4]),
		}
		forSecondNumber := map[bool]string{
			true:  strTicket[3:],
			false: string(strTicket[1]) + string(strTicket[3]) + string(strTicket[5]),
		}

		if getSum(forFirstNumber[mode == "Moscow"]) == getSum(forSecondNumber[mode == "Moscow"]) {
			luckyTickets = append(luckyTickets, strTicket)
		}
	}

	return
}

func getSum(strNumber string) (sum int) {
	for _, uCode := range strNumber {
		i, _ := strconv.Atoi(string(uCode))
		sum += i
	}
	return
}
