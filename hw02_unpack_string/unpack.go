package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) {

	if inputStr == "" {
		return inputStr, nil
	}

	runes := []rune(inputStr)
	lenRunes := len(runes) - 1
	var resultString strings.Builder

	for i := 0; i < len(runes[:lenRunes]); i++ {
		if unicode.IsDigit(runes[i+1]) {
			resultString.WriteString(strings.Repeat(string(runes[i]), int(runes[i+1])-48))
			i++
		} else {
			resultString.WriteRune(runes[i])
		}
	}

	if !unicode.IsDigit(runes[lenRunes]) {
		resultString.WriteRune(runes[len(runes)-1])
	}

	return resultString.String(), nil
}
