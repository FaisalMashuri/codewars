package __Kyu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCamelCaseMethod(t *testing.T) {
	testCase := []struct {
		input  string
		output string
	}{
		{input: "hello case", output: "HelloCase"},
		{input: "camel case word", output: "CamelCaseWord"},
		{input: "test case", output: "TestCase"},
		{input: "camel case method", output: "CamelCaseMethod"},
		{input: "", output: ""},
	}

	for _, val := range testCase {
		res := CamelCase(val.input)
		fmt.Println(res)
		assert.Equal(t, val.output, res)

	}
}

func CamelCase(s string) string {
	// your code here
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	words := strings.Split(s, " ")
	for i, word := range words {
		runeWord := []rune(word)
		stringRune := string(runeWord[0])
		stringRune = strings.ToUpper(stringRune)
		runeWord[0] = []rune(stringRune)[0]
		words[i] = string(runeWord)
	}

	return strings.Join(words, "")
}
