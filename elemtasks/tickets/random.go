package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func fillTickets() {
	var tickets string
	for i := 0; i < 100; i++ {
		tickets += getTicketNumber()
	}
	if err := writeFile(tickets); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Written successfully!")
}

func getTicketNumber() (ticket string) {
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		ticket += strconv.Itoa(rand.Intn(10))
		time.Sleep(time.Millisecond * 10)
	}
	ticket += "\n"
	return
}

func writeFile(tickets string) error {
	if err := ioutil.WriteFile("elemtasks/task6/tickets.txt", []byte(tickets), 0); err != nil {
		return err
	}
	return nil
}
