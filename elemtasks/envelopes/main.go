package main

import (
	"fmt"
	"strings"
)

type envelope struct {
	name  string
	side1 float64
	side2 float64
}

func main() {
	if err := runMainLoop(); err != nil {
		fmt.Println(err)
	}
}

func runMainLoop() error {
	for {
		envelopes := [2]*envelope{
			{name: "My envelope 1"},
			{name: "My envelope 2"},
		}

		for i := 0; i < 2; i++ {
			err := setEnvelope(envelopes[i])
			if err != nil {
				return err
			}
		}

		checkAnswer := check(envelopes[0], envelopes[1])
		fmt.Println(checkAnswer, "To continue: 'y' or 'yes') > ")

		userAnswer, err := getAnswerFromUser()
		if err != nil {
			return err
		}

		if isYes(userAnswer) {
			fmt.Println("Start new circle!")
			continue
		}
		break
	}
	return nil
}

func getAnswerFromUser() (answer string, err error) {
	_, err = fmt.Scan(&answer)
	if err != nil {
		return
	}
	return
}

func setEnvelope(env *envelope) error {
	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Printf("Input for <%s>. Side1: ", (*env).name)
			_, err := fmt.Scan(&(*env).side1)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Input for <%s>. Side2: ", env.name)
			_, err := fmt.Scan(&(*env).side2)
			if err != nil {
				return err
			}
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
