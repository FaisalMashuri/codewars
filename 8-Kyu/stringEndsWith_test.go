package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
Complete the solution so that it returns true
if the first argument(string) passed in ends with the 2nd argument (also a string).
*/

func TestStringEndsWith(t *testing.T) {
	testCase := []struct {
		input  []string
		output bool
	}{
		{
			input:  []string{"abc", "bc"},
			output: true,
		},
		{
			input:  []string{"abc", "d"},
			output: false,
		},
	}

	for _, val := range testCase {
		res := StringEndsWith(val.input[0], val.input[1])
		assert.Equal(t, val.output, res)
	}
	for _, val := range testCase {
		res2 := Solution2(val.input[0], val.input[1])
		assert.Equal(t, val.output, res2)
	}
}

func StringEndsWith(str, ending string) bool {
	res := strings.HasSuffix(str, ending)
	return res
}

func Solution2(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	endIndex := len(str) - 1
	for i := len(ending) - 1; i >= 0; i-- {
		if str[endIndex] != ending[i] {
			return false
		}
		endIndex--
	}
	return true
}
