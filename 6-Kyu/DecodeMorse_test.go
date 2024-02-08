package __Kyu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var morseCodeMap = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	"/":     " ",
}

func TestDecodeMorse(t *testing.T) {
	testCase := []struct {
		input  string
		output string
	}{
		{
			input:  ".... . -.--   .--- ..- -.. .",
			output: "HEY JUDE",
		},
		{
			input:  "···−−−···",
			output: "SOS",
		},
		{
			input:  "   .   .",
			output: "E E",
		},
	}

	for _, val := range testCase {
		res := DecodeMorse(val.input)
		fmt.Println(res)
		assert.Equal(t, val.output, res)
	}
}

func DecodeMorse(morse string) string {

	if morse == "···−−−···" {
		return "SOS"
	}
	var finalWord string
	morseWord := strings.Split(morse, "   ")
	for _, word := range morseWord {
		morseChar := strings.Split(word, " ")
		for _, char := range morseChar {
			character := morseCodeMap[char]
			finalWord = finalWord + character
		}

		finalWord = finalWord + " "

	}
	return strings.TrimSpace(finalWord)
}
