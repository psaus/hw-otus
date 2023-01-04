package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordCollector(t *testing.T) {
	t.Run("wordUp increase counter", func(t *testing.T) {
		collector := NewWordCollector()
		collector.WordUp("да")
		collector.WordUp("нет")
		collector.WordUp("да")
		require.Len(t, collector.list, 2)
		require.Equal(t, 2, collector.list["да"].counter)
	})
}
