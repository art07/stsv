// https://github.com/stretchr/testify
package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inBool  bool
	inStr   string
	inStr2  string
	inSlStr []string

	outErr   error
	outBool  bool
	outSlStr []string
	outInt   int
}

func TestFlagValidation(t *testing.T) {
	tests := []Test{
		{inBool: true, inStr: "./tickets.txt", inStr2: "Moscow", outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inStr: noData, inStr2: "Moscow", outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inStr: "./tickets.txt", inStr2: noData, outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inStr: "./tickets.txt", inStr2: "Mascaw", outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inStr: "./tickets.txt", inStr2: "PiterPen", outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inStr: "./tickets.txt", inStr2: "Moscow", outErr: nil},
		{inBool: false, inStr: "./tickets.txt", inStr2: "Piter", outErr: nil},
	}

	for _, test := range tests {
		err := flagValidation(test.inBool, test.inStr, test.inStr2)
		if err != nil {
			assert.NotNil(t, err)
			assert.EqualError(t, err, test.outErr.Error())
			continue
		}
		assert.Nil(t, err)
	}
}

func TestIsTicket(t *testing.T) {
	tests := []Test{
		{inStr: "123456", outBool: true},
		{inStr: "12345a", outBool: false},
		{inStr: "12345", outBool: false},
		{inStr: "1234567", outBool: false},
		{inStr: "zaqwsx", outBool: false},
	}

	for _, test := range tests {
		result, _ := isTicket(test.inStr)
		assert.EqualValues(t, test.outBool, result)
	}
}

func TestGetTickets(t *testing.T) {
	tests := []Test{
		{inStr: "./tickets_test.txt", outSlStr: []string{"901361", "123456", "040150"}},
		{inStr: "./tickets_test123.txt", outSlStr: []string{"901361", "123456", "040150"}},
	}

	for _, test := range tests {
		tickets, err := getTickets(test.inStr)
		if err != nil {
			assert.NotNil(t, err)
			assert.Nil(t, tickets)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, tickets)
			assert.EqualValues(t, test.outSlStr, tickets)
		}
	}
}

func TestGetSum(t *testing.T) {
	tests := []Test{
		{inStr: "123", outInt: 6},
		{inStr: "999", outInt: 27},
		{inStr: "000", outInt: 0},
	}

	for _, test := range tests {
		result := getSum(test.inStr)
		assert.EqualValues(t, test.outInt, result)
	}
}

func TestLuckyTickets(t *testing.T) {
	tests := []Test{
		{inSlStr: []string{"901361", "123456", "040150"}, inStr: "Moscow", outSlStr: []string{"901361"}},
		{inSlStr: []string{"901361", "123456", "040150"}, inStr: "Piter", outSlStr: []string{"040150"}},
	}

	for _, test := range tests {
		result := checkLuckyTickets(test.inSlStr, test.inStr)
		assert.EqualValues(t, test.outSlStr, result)
	}
}
