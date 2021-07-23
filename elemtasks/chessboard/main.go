package main

import (
	"errors"
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
	parseFlags(&help, &rows, &columns)

	if err := flagValidation(rows, columns, help); err != nil {
		fmt.Println(err, strForTask1)
		return
	}

	fmt.Println(getChessboard(rows, columns))
}

func parseFlags(help *bool, rows, columns *int) {
	flag.BoolVar(help, "h", false, "help")
	flag.IntVar(rows, "r", 1, "rows")
	flag.IntVar(columns, "c", 1, "columns")
	flag.Parse()
}

func flagValidation(rows, columns int, help bool) error {
	if rows < 1 || columns < 1 || help {
		return errors.New("interrupt the app because of flags")
	}
	return nil
}

func getChessboard(rows, columns int) string {
	var strBuild strings.Builder
	/*Беру ряд */
	for i := 1; i <= rows; i++ {
		/*и вывожу колонку.*/
		for i := 0; i < columns; i++ {
			strBuild.WriteString("* ")
		}

		/*Перехожу на новый ряд или с отступом или без.*/
		if i%2 != 0 {
			strBuild.WriteString("\n ")
		} else {
			strBuild.WriteString("\n")
		}
	}
	return strings.TrimSpace(strBuild.String())
}
