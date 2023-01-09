package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `qwe\\\3\\`, expected: `qwe\3\`},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `ab\n`, `qwe\\\3\`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestIsCanBeEscaped(t *testing.T) {
	validEscape := []rune{'\\', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	t.Run("valid runes for escaping", func(t *testing.T) {
		for _, symbol := range validEscape {
			require.True(t, isCanBeEscaped(symbol))
		}
	})

	notValidEscape := []rune{'a', 'd', '\n', '#', '?', '+', 'c', 'E', ' ', '_', '-'}
	t.Run("not valid runes for escaping", func(t *testing.T) {
		for _, symbol := range notValidEscape {
			require.False(t, isCanBeEscaped(symbol))
		}
	})
}
