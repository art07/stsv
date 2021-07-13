package main

import (
	"flag"
	"fmt"
)

var strForTask1 = "\ntask1 is a tool for printing chessboard.\n" +
	"Flags:\n" +
	"\tflag -r > number of rows (-r=4), default 1, min 1;\n" +
	"\tflag -c > number of columns (-c=8), default 1, min 1;\n" +
	"\tflag -h > help info (-h)."

var rows, columns int
var help bool

func main() {
	flag.BoolVar(&help, "h", false, "help")
	flag.IntVar(&rows, "r", 1, "rows")
	flag.IntVar(&columns, "c", 1, "columns")
	flag.Parse()

	if rows < 1 || columns < 1 || help {
		fmt.Println(strForTask1)
		return
	}

	/*Беру ряд */
	for i := 1; i <= rows; i++ {
		/*и вывожу колонку.*/
		for i := 0; i < columns; i++ {
			fmt.Printf("%s ", "*")
		}
		/*Перехожу на новый ряд или с отступом или без.*/
		if i%2 != 0 {
			fmt.Print("\n ")
		} else {
			fmt.Println()
		}
	}
}
