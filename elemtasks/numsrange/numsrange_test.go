// https://github.com/stretchr/testify
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inI  int
	inB  bool
	outS string
	outB bool
}

func TestGetRange(t *testing.T) {
	tests := []Test{
		{inI: 1, outS: ""},
		{inI: 5, outS: "1,2"},
		{inI: 150, outS: "1,2,3,4,5,6,7,8,9,10,11,12"},
	}

	for _, test := range tests {
		result := getRange(test.inI)
		assert.Equal(t, test.outS, result)
	}
}

func TestInterrupt(t *testing.T) {
	tests := []Test{
		{inB: true, inI: 1, outB: true},
		{inB: false, inI: 1, outB: false},
		{inB: true, inI: 0, outB: true},
		{inB: false, inI: 0, outB: true},
	}

	for _, test := range tests {
		result := interrupt(test.inB, test.inI)
		assert.Equal(t, test.outB, result)
	}
}
