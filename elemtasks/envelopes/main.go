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

var env1, env2 envelope

func main() {
	for {
		env1 = envelope{name: "My envelope 1"}
		env2 = envelope{name: "My envelope 2"}

		err := setEnvelope(&env1)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = setEnvelope(&env2)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(check())

		fmt.Print("To continue: 'y' or 'yes') > ")
		var answer string
		_, err = fmt.Scan(&answer)
		if err != nil {
			fmt.Println(err)
			break
		}

		if isYes(answer) {
			fmt.Println("Start new circle!")
		} else {
			break
		}
	}
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
