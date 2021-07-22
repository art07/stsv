// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	help       bool
	filePath   string
	searchWord string
	newWord    string
	wantTimes  int
	outB       bool
}

func TestNumberOfOccurrences(t *testing.T) {
	tests := []Test{
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "0123456789", wantTimes: 0},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "Lorem", wantTimes: 2},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text123.txt", searchWord: "Lorem", wantTimes: 0},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text123.txt", searchWord: "0123456789", wantTimes: 0},
	}

	for _, test := range tests {
		result, err := numberOfOccurrences(test.filePath, test.searchWord)
		if err != nil {
			assert.NotNil(t, err)
			assert.Equal(t, 0, result)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.wantTimes, result)
		}
	}
}
func TestChangeText(t *testing.T) {
	tests := []Test{
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "Lorem", newWord: "Lorem1"},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "0123456789", newWord: "9876543210"},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "Lorem1", newWord: "Lorem"},
		{filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text123.txt", searchWord: "Lorem", newWord: "Lorem1"},
	}

	for _, test := range tests {
		err := changeText(test.filePath, test.searchWord, test.newWord)
		if err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestInterrupt(t *testing.T) {
	tests := []Test{
		{help: true, filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "Lorem", outB: true},
		{help: false, filePath: noData, searchWord: "Lorem", outB: true},
		{help: false, filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: noData, outB: true},
		{help: false, filePath: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", searchWord: "Lorem", outB: false},
	}

	for _, test := range tests {
		result := interrupt(test.help, test.filePath, test.searchWord)
		assert.Equal(t, test.outB, result)
	}
}
