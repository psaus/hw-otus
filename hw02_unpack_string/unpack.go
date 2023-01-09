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
	inputRunes := []rune(input)
	inputLen := len(inputRunes)

	if inputLen == 0 {
		return "", nil
	}

	if unicode.IsDigit(inputRunes[0]) {
		return "", ErrInvalidString
	}

	var unpackedStringBuilder strings.Builder
	var index int

	for index < inputLen {
		if inputRunes[index] == EscapeRune {
			index++

			if index >= inputLen || !isCanBeEscaped(inputRunes[index]) {
				return "", ErrInvalidString
			}
		}

		currentRune := inputRunes[index]

		if isCanBeEscaped(currentRune) && inputRunes[index-1] != EscapeRune {
			return "", ErrInvalidString
		}

		nextRuneIndex := index + 1
		repeatRuneCount := 1

		if nextRuneIndex < inputLen && unicode.IsDigit(inputRunes[nextRuneIndex]) {
			repeatRuneCount, _ = strconv.Atoi(string(inputRunes[nextRuneIndex]))
			index++
		}

		if repeatRuneCount > 0 {
			unpackedStringBuilder.WriteString(strings.Repeat(string(currentRune), repeatRuneCount))
		}

		index++
	}

	return unpackedStringBuilder.String(), nil
}
