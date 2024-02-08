package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInverValues(t *testing.T) {
	testCase := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{-1, -2, -3, -4, -5},
		},
		{
			input:  []int{1, -2, 3, -4, 5},
			output: []int{-1, 2, -3, 4, -5},
		},
	}
	for _, value := range testCase {
		res := InverValues(value.input)
		assert.Equal(t, value.output, res)
	}
}

func InverValues(arr []int) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = arr[i] * -1
	}
	return result
}
