package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

var strForTask5 = "\ntask5 is a tool to convert numbers.\n" +
	"Convert numbers to string view. Max value => 100000.\n" +
	"Args:\n" +
	"\targ1 number (MAX> 100000)."

var numbersMap = map[int]string{
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
	10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
	17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	/*------------*/
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
}

func main() {
	/*Если не 2 аргумента strForTask5.*/
	if len(os.Args) != 2 {
		fmt.Println(strForTask5)
		return
	}

	/*Получаю число из аргумента. Если ошибка, strForTask5.*/
	number, err := getNumber()
	if err != nil {
		fmt.Println(strForTask5)
		return
	}

	switch {
	case number > 100_000:
		fmt.Println("too much for me!")
	case number == 100_000:
		fmt.Println("one hundred thousand")
	default:
		fmt.Println(recursiveFunc(number))
	}
}

/*1-99*/
func range1To99(number int) (textNumber string) {
	// 1-20 просто достаю из map.
	if number <= 20 {
		textNumber = numbersMap[number]
		return
	}

	// Если кратное 10, достаю из map, в остальном собираю строку.
	if rest := number % 10; rest == 0 {
		textNumber = numbersMap[number]
	} else {
		textNumber = fmt.Sprintf("%s %s", numbersMap[number-rest], numbersMap[rest])
	}
	return
}

/*100-99_999*/
func recursiveFunc(numberToConvert int) (textNumber string) {
	/*Получаю ключевое слово в зависимости от колличества цифр.*/
	mainTextPart := getTextPart(numberToConvert)

	/*Если число равное 2 цифрам, просто обрабатываю в range1To99 и выхожу из recursiveFunc.*/
	if mainTextPart == "" {
		textNumber = range1To99(numberToConvert)
		return
	}

	/*Получение делителя для расчетов в зависимости от колличества цифр.*/
	divider := map[string]int{"thousand": 1000, "hundred": 100}[mainTextPart]

	/*Получаю колличество сотен/тысяч.*/
	hundredsThousands := numberToConvert / divider

	/*Получение остатка для последующей обработки.*/
	rest := numberToConvert % divider

	/*Если сотен/тысяч до 10, просто формирую стринг с числом; иначе просто обрабатываю в range1To99.*/
	if hundredsThousands <= 9 {
		textNumber = fmt.Sprintf("%s %s", numbersMap[hundredsThousands], mainTextPart)
	} else {
		textNumber = fmt.Sprintf("%s %s", range1To99(hundredsThousands), mainTextPart)
	}

	/*Если rest = 0, дальнейшая обработка не требуется;*/
	if rest != 0 {
		if rest < 100 {
			/*если rest меньше 100, обрабатываю в range1To9 и выход;*/
			textNumber = fmt.Sprintf("%s %s", textNumber, range1To99(rest))
		} else {
			/*!!! если rest 100 и больше, необходимо рекурсивно вызвать эту же функцию для дальнейшей
			обработки числа из rest !!!*/
			textNumber = fmt.Sprintf("%s %s", textNumber, recursiveFunc(rest))
		}
	}
	return
}

/*Попытка получить число из аргумента.*/
func getNumber() (int, error) {
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, err
	}
	return i, nil
}

/*Возврат текстовой метки в зависимости числа.*/
func getTextPart(i int) (mainTextPart string) {
	switch utf8.RuneCountInString(strconv.Itoa(i)) {
	case 5:
		fallthrough
	case 4:
		mainTextPart = "thousand"
	case 3:
		mainTextPart = "hundred"
	case 2:
		mainTextPart = ""
	}
	return
}
