package main

import (
	"fmt"
	"strings"
)

var input = [4]string{"a", "b", "c", "d"}
var numbers [4]float64 // (numbers[0]=a,numbers[1]=b) (numbers[2]=c,numbers[3]=d)

func main() {
	for {
		err := fillNumbers()
		if err != nil {
			fmt.Println("Repeat...")
			continue
		}

		if numbers[0] > numbers[2] && numbers[1] > numbers[3] {
			fmt.Println("An envelope (c,d) can be inserted into an envelope (a,b) ")
		} else if numbers[0] < numbers[2] && numbers[1] < numbers[3] {
			fmt.Println("An envelope (a,b) can be inserted into an envelope (c,d) ")
		} else {
			fmt.Println("Envelopes cannot be nested.")
		}

		fmt.Print("To continue: 'y' or 'yes') > ")
		var answer string
		_, _ = fmt.Scan(&answer)

		if isYes(answer) {
			fmt.Println("Start new circle!")
		} else {
			fmt.Println("Program finished!")
			return
		}
	}
}

func fillNumbers() error {
	for i := 0; i < 4; i++ {
		fmt.Printf("Input for <%s>: ", input[i])
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("Wrong input")
			return err
		}
	}
	return nil
}

func isYes(answer string) (b bool) {
	answer = strings.ToLower(answer)
	if answer == "yes" || answer == "y" {
		b = true
	}
	return
}
