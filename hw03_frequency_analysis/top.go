package hw03frequencyanalysis

import (
	"regexp"
	"strings"
)

var cleanUpWordRegexp = regexp.MustCompile(`[^a-zA-Z\\u0430-яА-Я-]+`)

func Top10(text string) []string {
	collector := NewWordCollector()

	for _, value := range strings.Fields(text) {
		value = strings.Trim(value, "-")
		if len(value) == 0 {
			continue
		}

		value := cleanUpWordRegexp.ReplaceAllString(value, "")
		collector.WordUp(strings.ToLower(value))
	}

	return collector.GetWordFrenquencySorter().GetSortedWords(10)
}
