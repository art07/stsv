// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	In  string
	In2 string
	Out int
}

func TestNumberOfOccurrences(t *testing.T) {
	tests := []Test{
		{In: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", In2: "0123456789", Out: 0},
		{In: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text.txt", In2: "Lorem", Out: 2},
		{In: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text123.txt", In2: "Lorem", Out: 0},
		{In: "D:\\IT\\Golang\\go\\src\\art\\edu\\stsv\\elemtasks\\parser\\text123.txt", In2: "0123456789", Out: 0},
	}

	for _, test := range tests {
		result, err := numberOfOccurrences(test.In, test.In2)
		if err != nil {
			assert.NotNil(t, err)
			assert.Equal(t, 0, result)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.Out, result)
		}
	}
}
func TestChangeText(t *testing.T) {

}
