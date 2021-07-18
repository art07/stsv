// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

type testData struct {
	list []string
	want string
}

func TestPrepareAnswer(t *testing.T) {
	tests := []testData{
		{list: []string{}, want: "[ mode]. Lucky tickets in file = [0]\n"},
		{list: []string{"003030"}, want: "[ mode]. Lucky tickets in file = [1]\nLucky array > [003030]\n"},
		{list: []string{"003030", "030003", "300300"}, want: "[ mode]. Lucky tickets in file = [3]\nLucky array > [003030 030003 300300]\n"},
	}

	for _, test := range tests {
		result := prepareAnswer(test.list)
		// assert equality
		assert.Equal(t, test.want, result, "they should be equal")
	}
}

func TestGetTicketNumber(t *testing.T) {
	result := getTicketNumber()
	// assert equality
	assert.Equal(t, 6, utf8.RuneCountInString(result), "they should be equal")
}

func TestWriteFile(t *testing.T) {
	err := writeFile("901361\n123456\n040150", "D:/IT/Golang/go/src/art/edu/stsv/elemtasks/tickets/tickets_test.txt")
	// assert for nil
	assert.Nil(t, err)
}

func TestGetTickets(t *testing.T) {
	tickets, err := getTickets("D:/IT/Golang/go/src/art/edu/stsv/elemtasks/tickets/tickets_test.txt")
	// assert equality
	assert.Equal(t, []string{"901361", "123456", "040150"}, tickets, "they should be equal")
	// assert for nil
	assert.Nil(t, err)
}

func TestFillTickets(t *testing.T) {
	err := fillTickets("D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\tickets\\tickets.txt")
	// assert for nil
	assert.Nil(t, err)
}

type testData2 struct {
	td     testData
	mode   string
	wantSl []string
}

func TestLuckyTickets(t *testing.T) {
	tests := []testData2{
		{td: testData{list: []string{}}, mode: "Moscow", wantSl: []string{}},
		{td: testData{list: []string{}}, mode: "Piter", wantSl: []string{}},
		{td: testData{list: []string{"901361", "123456", "040150"}}, mode: "Moscow", wantSl: []string{"901361"}},
		{td: testData{list: []string{"901361", "123456", "040150"}}, mode: "Piter", wantSl: []string{"040150"}},
		{td: testData{list: []string{"123456", "234567", "345678"}}, mode: "Moscow", wantSl: []string{}},
		{td: testData{list: []string{"123456", "234567", "345678"}}, mode: "Piter", wantSl: []string{}},
		{td: testData{list: []string{"901361", "123456", "040150", "912660", "011000"}}, mode: "Moscow", wantSl: []string{"901361", "912660"}},
		{td: testData{list: []string{"901361", "123456", "040150", "912660", "011000"}}, mode: "Piter", wantSl: []string{"040150", "011000"}},
	}

	for _, test := range tests {
		result, _ := checkLuckyTickets(test.td.list, test.mode)
		// assert equality
		assert.Equal(t, test.wantSl, result, "they should be equal")
	}
}

type testData3 struct {
	strNumber string
	want      int
}

func TestGetSum(t *testing.T) {
	tests := []testData3{
		{"123", 6},
	}

	for _, test := range tests {
		result, _ := getSum(test.strNumber)
		// assert equality
		assert.Equal(t, test.want, result, "they should be equal")
	}
}
