package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type triangle struct {
	Name  string
	Side1 float64
	Side2 float64
	Side3 float64
	Area  float64
}

var triangles = make([]*triangle, 0, 8)

const header = "============= Triangles list: ===============\n"

func main() {
	for {
		var input string
		fmt.Print("<имя>,<длина стороны>,<длина стороны>,<длина стороны>: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dataArr := strings.Split(input, ",")

		if len(dataArr) != 4 {
			fmt.Println("Not enough data")
			continue
		}

		tnl, err := createTriangle(dataArr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		triangles = append(triangles, tnl)

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

func createTriangle(arr []string) (*triangle, error) {
	var side1, side2, side3 float64
	name := arr[0]
	side1, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return nil, err
	}
	side2, err = strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return nil, err
	}
	side3, err = strconv.ParseFloat(arr[3], 64)
	if err != nil {
		return nil, err
	}

	tnl := &triangle{
		Name:  name,
		Side1: side1,
		Side2: side2,
		Side3: side3,
	}

	ok := isTriangle(tnl)
	if !ok {
		return nil, errors.New("figure not triangle")
	}

	p := (side1 + side2 + side3) / 2
	tnl.Area = math.Sqrt(p * (p - side1) * (p - side2) * (p - side3))

	return tnl, nil
}

func isTriangle(t *triangle) (b bool) {
	if t.Side1+t.Side2 > t.Side3 && t.Side1+t.Side3 > t.Side2 && t.Side2+t.Side3 > t.Side1 {
		b = true
	}
	return
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
