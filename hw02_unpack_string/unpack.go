package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const unicodeZero = 0x30
const emptyString = ""

func Unpack(inputStr string) (string, error) {

	if inputStr == emptyString {
		return emptyString, nil
	}

	runes := []rune(inputStr)
	lenRunes := len(runes) - 1
	var resultString strings.Builder

	if unicode.IsDigit(runes[0]) {
		return emptyString, ErrInvalidString
	}

	for i := 0; i < len(runes[:lenRunes]); i++ {
		if unicode.IsDigit(runes[i+1]) {
			resultString.WriteString(strings.Repeat(string(runes[i]), int(runes[i+1])-unicodeZero))
			i++
			if i == lenRunes { //last rune is digit
				break
			}
			err := check2Digits(runes[i+1])
			if err != nil {
				return emptyString, err
			}
		} else {
			resultString.WriteRune(runes[i])
		}
	}

	if !unicode.IsDigit(runes[lenRunes]) {
		resultString.WriteRune(runes[len(runes)-1])
	}

	return resultString.String(), nil
}

func check2Digits(a rune) error {
	if unicode.IsDigit(a) {
		return ErrInvalidString
	}
	return nil
}
