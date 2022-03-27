package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Validate(char rune) bool {
	if char == 0 || unicode.IsDigit(char) {
		return false
	}
	return true
}

func Unpack(inputString string) (string, error) {
	if len(inputString) == 0 {
		return "", nil
	}
	var outputString strings.Builder
	var previousChar rune
	var repeatCount int

	for _, char := range inputString {
		if unicode.IsDigit(char) {
			if !Validate(previousChar) {
				return "", ErrInvalidString
			}
			repeatCount, _ = strconv.Atoi(string(char))
			outputString.WriteString(strings.Repeat(string(previousChar), repeatCount))
		} else if Validate(previousChar) {
			outputString.WriteRune(previousChar)
		}
		previousChar = char
	}

	if Validate(previousChar) {
		outputString.WriteRune(previousChar)
	}

	return outputString.String(), nil
}
