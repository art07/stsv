// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inStr  string
	inStr2 string
	inStr3 string
	inBool bool

	outInt int
	outStr string
}

func TestFlagValidation(t *testing.T) {
	tests := []Test{
		{inBool: true, inStr: "./text.txt", inStr2: "Lorem"},  // error
		{inBool: false, inStr: noData, inStr2: "Lorem"},       // error
		{inBool: false, inStr: "./text.txt", inStr2: noData},  // error
		{inBool: false, inStr: "./text.txt", inStr2: "Lorem"}, // no error
	}

	for _, test := range tests {
		err := flagValidation(test.inBool, test.inStr, test.inStr2)
		if err != nil {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
	}
}

func TestMode1(t *testing.T) {
	tests := []Test{
		{
			inStr:  "./text1.txt",
			inStr2: "Lorem",
			outStr: "open ./text1.txt: The system cannot find the file specified.",
		},
		{
			inStr:  "./text.txt",
			inStr2: "Lorem",
			outStr: "Mode #1 finished successfully! [Lorem] meets = 2 time(s).",
		},
	}

	for _, test := range tests {
		result := mode1(test.inStr, test.inStr2)
		assert.EqualValues(t, test.outStr, result)
	}
}

func TestMode2(t *testing.T) {
	tests := []Test{
		{
			inStr:  "./text1.txt",
			inStr2: "Lorem",
			inStr3: "Lorem123",
			outStr: "open ./text1.txt: The system cannot find the file specified.",
		},
		{
			inStr:  "./text.txt",
			inStr2: "Lorem",
			inStr3: "Lorem123",
			outStr: "Mode #2 finished successfully! [Lorem123] meets = 2 time(s).",
		},
		{
			inStr:  "./text.txt",
			inStr2: "Lorem123",
			inStr3: "Lorem",
			outStr: "Mode #2 finished successfully! [Lorem] meets = 2 time(s).",
		},
	}

	for _, test := range tests {
		result := mode2(test.inStr, test.inStr2, test.inStr3)
		assert.EqualValues(t, test.outStr, result)
	}
}

func TestNumberOfOccurrences(t *testing.T) {
	tests := []Test{
		{inStr: "./text.txt", inStr2: "Lorem", outInt: 2},
		{inStr: "", inStr2: "Lorem", outInt: 0},
		{inStr: "./text.txt", inStr2: "0123456789", outInt: 0},
		{inStr: "./text.txt", inStr2: "viverra", outInt: 3},
	}

	for _, test := range tests {
		result, err := numberOfOccurrences(test.inStr, test.inStr2)
		if err != nil {
			assert.NotNil(t, err)
			assert.EqualValues(t, test.outInt, result)
		} else {
			assert.Nil(t, err)
			assert.EqualValues(t, test.outInt, result)
		}
	}
}

func TestChangeText(t *testing.T) {
	tests := []Test{
		{inStr: "./text.txt", inStr2: "Lorem", inStr3: "Lorem123", outInt: 2},
		{inStr: "./text.txt", inStr2: "Lorem123", inStr3: "Lorem", outInt: 2},
		{inStr: "./text123.txt", inStr2: "Lorem", inStr3: "Lorem123", outInt: 0},
		{inStr: "./text.txt", inStr2: "0123456789", inStr3: "Lorem123", outInt: 0},
	}

	for _, test := range tests {
		result, err := changeText(test.inStr, test.inStr2, test.inStr3)
		if err != nil {
			assert.NotNil(t, err)
			assert.EqualValues(t, test.outInt, result)
		} else {
			assert.Nil(t, err)
			assert.EqualValues(t, test.outInt, result)
		}
	}
}
