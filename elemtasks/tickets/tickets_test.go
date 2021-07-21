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

type InOut struct {
	in  string
	in2 string
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
	tests := [2]InOut{
		{in: "901361\n123456\n040150", in2: "D:/IT/Golang/go/src/art/edu/stsv/elemtasks/tickets/tickets_test.txt"},
	}

	for _, test := range tests {
		err := writeFile(test.in, test.in2)
		if err != nil {
			// assert for not nil
			assert.NotNil(t, err)
		} else {
			// assert for nil
			assert.Nil(t, err)
		}
	}
}

func TestGetTickets(t *testing.T) {
	tests := [2]InOut{
		{in: "D:/IT/Golang/go/src/art/edu/stsv/elemtasks/tickets/tickets_test.txt"},
		{in: "D:/IT/Golang/go/src/art/edu/stsv/elemtasks/tickets/tickets_test_123.txt"},
	}

	for _, test := range tests {
		tickets, err := getTickets(test.in)
		if err != nil {
			assert.NotNil(t, err)
			assert.Nil(t, tickets)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, tickets)
			assert.Equal(t, []string{"901361", "123456", "040150"}, tickets, "they should be equal")
		}
	}
}

func TestFillTickets(t *testing.T) {
	tests := [2]InOut{
		{in: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\tickets\\tickets.txt"},
	}

	for _, test := range tests {
		err := fillTickets(test.in)
		if err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
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
		{td: testData{list: []string{"123def"}}, mode: "Moscow", wantSl: []string{"040150", "011000"}},
		{td: testData{list: []string{"abcdef"}}, mode: "Piter", wantSl: []string{"040150", "011000"}},
	}

	for _, test := range tests {
		result, err := checkLuckyTickets(test.td.list, test.mode)
		if err != nil {
			assert.NotNil(t, err)
			assert.Nil(t, result)
			assert.NotEqual(t, test.wantSl, result)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, test.wantSl, result, "they should be equal")
		}
	}
}

type testData3 struct {
	strNumber string
	want      int
}

func TestGetSum(t *testing.T) {
	tests := []testData3{
		{"123", 6},
		{"abc", 6},
	}

	for _, test := range tests {
		result, err := getSum(test.strNumber)
		if err != nil {
			assert.NotNil(t, err)
			assert.Equal(t, 0, result)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.want, result, "they should be equal")
		}
	}
}
