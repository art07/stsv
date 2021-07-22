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
	outB bool
	outS string
}

func TestInterrupt(t *testing.T) {
	tests := []Test{
		{inI: 0, inI2: 1, inB: false, outB: true},
		{inI: 1, inI2: 0, inB: false, outB: true},
		{inI: 1, inI2: 1, inB: true, outB: true},
		{inI: 1, inI2: 1, inB: false, outB: false},
	}

	for _, test := range tests {
		result := interrupt(test.inI, test.inI2, test.inB)
		assert.Equal(t, test.outB, result)
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
		assert.Equal(t, test.outS, result)
	}
}
