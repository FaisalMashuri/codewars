package __Kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Count the number of divisors of a positive integer n.

Random tests go up to n = 500000.

Examples (input --> output)
4 --> 3 // we have 3 divisors - 1, 2 and 4
5 --> 2 // we have 2 divisors - 1 and 5
12 --> 6 // we have 6 divisors - 1, 2, 3, 4, 6 and 12
30 --> 8 // we have 8 divisors - 1, 2, 3, 5, 6, 10, 15 and 30
Note you should only return a number, the count of divisors.
The numbers between parentheses are shown only for you to see which numbers are counted in each case.
*/

func TestCountDivisor(t *testing.T) {
	testCase := []struct {
		input  int
		output int
	}{
		{
			input:  4,
			output: 3,
		},
		{
			input:  5,
			output: 2,
		},
		{input: 12, output: 6},
		{input: 30, output: 8},
	}

	for _, val := range testCase {
		res := Divisors(val.input)
		assert.Equal(t, val.output, res)
	}
}

func Divisors(n int) int {
	if n > 500000 {
		return 0
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
