// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inI  int
	inI2 int
	inB  bool
	outS string
}

func TestFlagValidation(t *testing.T) {
	tests := []Test{
		{inI: 0, inI2: 1, inB: false},
		{inI: 1, inI2: 0, inB: false},
		{inI: 1, inI2: 1, inB: true},
		{inI: 1, inI2: 1, inB: false},
	}

	for _, test := range tests {
		err := flagValidation(test.inI, test.inI2, test.inB)
		if err != nil {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
	}
}

func TestGetChessboard(t *testing.T) {
	tests := []Test{
		{inI: 1, inI2: 1, outS: "*"},
		{inI: 2, inI2: 1, outS: "* \n *"},
		{inI: 1, inI2: 2, outS: "* *"},
		{inI: 2, inI2: 2, outS: "* * \n * *"},
	}

	for _, test := range tests {
		result := getChessboard(test.inI, test.inI2)
		assert.EqualValues(t, test.outS, result)
	}
}
