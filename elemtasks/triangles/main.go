package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var triangles = make([]*triangle, 0, 8)
var input string
var dataArr []string

const header = "\n============= Triangles list: ===============\n"

func main() {
	for {
		/*Ввод csv строки.*/
		fmt.Print("<имя>,<длина стороны>,<длина стороны>,<длина стороны>: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			dataArr = strings.Split(input, ",")
			if len(dataArr) != 4 {
				fmt.Println("Not enough data")
				continue
			}
		}

		/*Создать треугольник.*/
		tnl, err := createTriangle(dataArr)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			triangles = append(triangles, tnl)
		}

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
			fmt.Println(prepareAnswer())
			fmt.Println("Program finished!")
			break
		}
	}
}

/*Попытка создать треугольник.*/
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

	ok := tnl.isTriangle()
	if !ok {
		return nil, errors.New("figure not triangle")
	} else {
		tnl.setHeronArea()
	}
	return tnl, nil
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
		s += fmt.Sprintf("%d. %s\n", i+1, t.getString())
	}
	return
}
