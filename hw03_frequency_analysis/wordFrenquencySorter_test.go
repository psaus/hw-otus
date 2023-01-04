package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordFrenquencySorter(t *testing.T) {
	dataSet := []string{"наверное", "нет", "Да", "да"}
	ws := &wordFrenquencySorter{
		wordStats: []wordFrenquency{
			{counter: 1, word: "да"},
			{counter: 1, word: "Да"},
			{counter: 3, word: "нет"},
			{counter: 6, word: "наверное"},
		},
	}

	t.Run("check sort", func(t *testing.T) {
		for i, v := range ws.GetSortedWords(0) {
			require.Equal(t, dataSet[i], v)
		}
	})

	t.Run("check slicing", func(t *testing.T) {
		require.Len(t, ws.GetSortedWords(2), 2)
	})
}
