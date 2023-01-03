package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const EscapeRune rune = '\\'

func isCanBeEscaped(r rune) bool {
	return EscapeRune == r || unicode.IsDigit(r)
}

func Unpack(input string) (string, error) {

	runes := []rune(input)
	inputLen := len(runes)

	if inputLen == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var b strings.Builder
	var index int = 0

	for index < inputLen {
		if runes[index] == EscapeRune {
			index++
		}

		currentRune := runes[index]

		if isCanBeEscaped(currentRune) && runes[index-1] != EscapeRune {
			return "", ErrInvalidString
		}

		nextRuneIndex := index + 1
		p := 1

		if nextRuneIndex < inputLen && unicode.IsDigit(runes[nextRuneIndex]) {
			p, _ = strconv.Atoi(string(runes[nextRuneIndex]))
			index++
		}

		for p > 0 {
			b.WriteRune(currentRune)
			p--
		}

		index++
	}

	return b.String(), nil
}
