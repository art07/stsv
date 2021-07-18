package main

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func fillTickets(filePath string) error {
	var strBuildTickets strings.Builder
	for i := 0; i < 100; i++ {
		strBuildTickets.WriteString(getTicketNumber() + "\n")
	}

	if err := writeFile(strBuildTickets.String(), filePath); err != nil {
		return err
	}
	return nil
}

func getTicketNumber() string {
	var strBuild strings.Builder
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		strBuild.WriteString(strconv.Itoa(rand.Intn(10)))
		time.Sleep(time.Millisecond * 3)
	}
	return strBuild.String()
}

func writeFile(tickets, filePath string) error {
	if err := ioutil.WriteFile(filePath, []byte(tickets), 0); err != nil {
		return err
	}
	return nil
}
