// https://github.com/stretchr/testify
package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inBool bool
	inInt  int

	outErr error
	outStr string
}

func TestFlagValidation(t *testing.T) {
	tests := []Test{
		{inBool: true, inInt: 1, outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inInt: 0, outErr: errors.New("interrupt the app because of flags")},
		{inBool: false, inInt: 1, outErr: nil},
	}

	for _, test := range tests {
		err := flagValidation(test.inBool, test.inInt)
		if err != nil {
			assert.NotNil(t, err)
			assert.EqualError(t, err, test.outErr.Error())
			continue
		}
		assert.Nil(t, err)
	}
}

func TestGetRange(t *testing.T) {
	tests := []Test{
		{inInt: 1, outStr: ""},
		{inInt: 5, outStr: "1,2"},
		{inInt: 150, outStr: "1,2,3,4,5,6,7,8,9,10,11,12"},
	}

	for _, test := range tests {
		result := getRange(test.inInt)
		assert.EqualValues(t, test.outStr, result)
	}
}
