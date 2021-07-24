package randtick

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func FillTickets(filePath string, ticketsNumber int) error {
	var strBuild strings.Builder

	for i := 0; i < ticketsNumber; i++ {
		strBuild.WriteString(getTicket())
	}

	if err := writeFile(strBuild.String(), filePath); err != nil {
		return err
	}
	return nil
}

func getTicket() string {
	var strBuild strings.Builder
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		strBuild.WriteString(strconv.Itoa(rand.Intn(10)))
		time.Sleep(time.Millisecond * 3)
	}
	return strBuild.String() + "\n"
}

func writeFile(tickets, filePath string) error {
	if err := os.WriteFile(filePath, []byte(tickets), 0); err != nil {
		return err
	}
	return nil
}
