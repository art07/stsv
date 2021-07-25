package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inStr  string
	inEnv1 *envelope
	inEnv2 *envelope

	outBool bool
	outStr  string
}

func TestIsYes(t *testing.T) {
	tests := []Test{
		{inStr: "123", outBool: false},
		{inStr: "yes", outBool: true},
		{inStr: "y", outBool: true},
		{inStr: "ye", outBool: false},
	}

	for _, test := range tests {
		result := isYes(test.inStr)
		assert.EqualValues(t, test.outBool, result)
	}
}

func TestStep1(t *testing.T) {
	tests := []Test{
		{
			inEnv1:  &envelope{name: "My envelope 1", side1: 3, side2: 7},
			inEnv2:  &envelope{name: "My envelope 2", side1: 5, side2: 20},
			outBool: false,
		},
		{
			inEnv1:  &envelope{name: "My envelope 1", side1: 15, side2: 30},
			inEnv2:  &envelope{name: "My envelope 2", side1: 5, side2: 10},
			outBool: true,
		},
	}

	for _, test := range tests {
		result := step1(test.inEnv1, test.inEnv2)
		assert.EqualValues(t, test.outBool, result)
	}
}

func TestStep2(t *testing.T) {
	tests := []Test{
		{
			inEnv1:  &envelope{name: "My envelope 1", side1: 5, side2: 10},
			inEnv2:  &envelope{name: "My envelope 2", side1: 15, side2: 30},
			outBool: false,
		},
		{
			inEnv1:  &envelope{name: "My envelope 1", side1: 15, side2: 30},
			inEnv2:  &envelope{name: "My envelope 2", side1: 5, side2: 10},
			outBool: true,
		},
	}

	for _, test := range tests {
		result := step2(test.inEnv1, test.inEnv2)
		assert.EqualValues(t, test.outBool, result)
	}
}

func TestCheck(t *testing.T) {
	tests := []Test{
		{
			inEnv1: &envelope{name: "My envelope 1", side1: 15, side2: 30},
			inEnv2: &envelope{name: "My envelope 2", side1: 7, side2: 14},
			outStr: "[My envelope 2] может поместиться в [My envelope 1].",
		},
		{
			inEnv1: &envelope{name: "My envelope 1", side1: 15, side2: 30},
			inEnv2: &envelope{name: "My envelope 2", side1: 7, side2: 15},
			outStr: "Поворот помог и теперь [My envelope 2] таки можно вложить в [My envelope 1].",
		},
		{
			inEnv1: &envelope{name: "My envelope 1", side1: 5, side2: 10},
			inEnv2: &envelope{name: "My envelope 2", side1: 11, side2: 20},
			outStr: "[My envelope 1] может поместиться в [My envelope 2].",
		},
		{
			inEnv1: &envelope{name: "My envelope 1", side1: 10, side2: 20},
			inEnv2: &envelope{name: "My envelope 2", side1: 18, side2: 30},
			outStr: "Теперь [My envelope 1] может поместиться в [My envelope 2].",
		},
		{
			inEnv1: &envelope{name: "My envelope 1", side1: 10, side2: 10},
			inEnv2: &envelope{name: "My envelope 2", side1: 10, side2: 10},
			outStr: "Неудача!) Ни [My envelope 1] не может быть вложен в [My envelope 2], ни [My envelope 2] не может быть вложен в [My envelope 1].",
		},
	}

	for _, test := range tests {
		result := check(test.inEnv1, test.inEnv2)
		assert.EqualValues(t, test.outStr, result)
	}
}
