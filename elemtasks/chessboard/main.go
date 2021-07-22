package main

import (
	"flag"
	"fmt"
	"strings"
)

var strForTask1 = "\ntask1 is a tool for printing chessboard.\n" +
	"Flags:\n" +
	"\tflag -r > number of rows (-r=4), default 1, min 1;\n" +
	"\tflag -c > number of columns (-c=8), default 1, min 1;\n" +
	"\tflag -h > help info (-h)."

func main() {
	var rows, columns int
	var help bool
	flag.BoolVar(&help, "h", false, "help")
	flag.IntVar(&rows, "r", 1, "rows")
	flag.IntVar(&columns, "c", 1, "columns")
	flag.Parse()

	if ok := interrupt(rows, columns, help); ok {
		return
	}

	fmt.Println(getChessboard(rows, columns))
}

//goland:noinspection GoUnhandledErrorResult
func getChessboard(rows, columns int) string {
	var strBuild strings.Builder
	/*Беру ряд */
	for i := 1; i <= rows; i++ {
		/*и вывожу колонку.*/
		for i := 0; i < columns; i++ {
			fmt.Fprintf(&strBuild, "%s ", "*")
		}

		/*Перехожу на новый ряд или с отступом или без.*/
		if i%2 != 0 {
			fmt.Fprint(&strBuild, "\n ")
		} else {
			fmt.Fprint(&strBuild, "\n")
		}
	}
	return strings.TrimSpace(strBuild.String())
}

func interrupt(rows, columns int, help bool) bool {
	if rows < 1 || columns < 1 || help {
		fmt.Println(strForTask1)
		return true
	}
	return false
}
