package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type triangle struct {
	Name string
	Area float64
}

var triangles = make([]triangle, 0, 8)

const header = "============= Triangles list: ===============\n"

func main() {
	for {
		fmt.Print("<имя>, <длина стороны>, <длина стороны>, <длина стороны>: ")
		snr := bufio.NewScanner(os.Stdin)
		snr.Scan()
		dataArr := strings.Split(snr.Text(), ",")
		if len(dataArr) != 4 {
			fmt.Println("Wrong input")
			continue
		}
		err := setNewDataToStruct(dataArr)
		if err != nil {
			fmt.Println("setNewDataToStruct failed")
			continue
		}

		fmt.Print("To continue: 'y' or 'yes') > ")
		var answer string
		_, _ = fmt.Scan(&answer)

		if isYes(answer) {
			fmt.Println("Start new circle!")
		} else {
			fmt.Println(prepareAnswer())
			fmt.Println("Program finished!")
			return
		}
	}
}

func setNewDataToStruct(arr []string) error {
	var side1, side2, side3, area float64
	name := arr[0]
	side1, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return err
	}
	side2, err = strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return err
	}
	side3, err = strconv.ParseFloat(arr[3], 64)
	if err != nil {
		return err
	}

	p := (side1 + side2 + side3) / 2
	area = math.Sqrt(p * (p - side1) * (p - side2) * (p - side3))

	triangles = append(triangles, triangle{Name: name, Area: area})
	return nil
}

func isYes(answer string) (b bool) {
	answer = strings.ToLower(answer)
	if answer == "yes" || answer == "y" {
		b = true
	}
	return
}

func prepareAnswer() (s string) {
	sort.Slice(triangles, func(i, j int) bool {
		return triangles[i].Area > triangles[j].Area
	})

	for i, t := range triangles {
		if i == 0 {
			s += header
		}
		if t.Area != 0.00 {
			s += fmt.Sprintf("%d. [%s]: %.2f сm\n", i+1, t.Name, t.Area)
		} else {
			s += fmt.Sprintf("%d. [%s]: sides of a triangle entered incorrectly!\n", i+1, t.Name)
		}
	}
	return
}
