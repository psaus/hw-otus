package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordFrenquency struct {
	counter int
	word    string
}

var cleanUpWordRegexp = regexp.MustCompile(`[^a-zA-Z\\u0430-яА-Я-]+`)

func convertWordCounterMapToSlice(wordCounterMap *map[string]int) []wordFrenquency {
	wordFrenquencyList := []wordFrenquency{}
	for w, c := range *wordCounterMap {
		wordFrenquencyList = append(wordFrenquencyList, wordFrenquency{c, w})
	}
	return wordFrenquencyList
}

func Top10(text string) []string {
	wordCounterMap := make(map[string]int)

	for _, value := range strings.Fields(text) {
		value = strings.Trim(value, " -")
		if len(value) == 0 {
			continue
		}

		value := strings.ToLower(cleanUpWordRegexp.ReplaceAllString(value, ""))
		wordCounterMap[value]++
	}

	wordFrenquencyList := convertWordCounterMapToSlice(&wordCounterMap)

	sort.Slice(wordFrenquencyList, func(i, j int) bool {
		if wordFrenquencyList[i].counter == wordFrenquencyList[j].counter {
			return wordFrenquencyList[i].word < wordFrenquencyList[j].word
		}

		return wordFrenquencyList[i].counter > wordFrenquencyList[j].counter
	})

	returnSize := 10
	if returnSize > len(wordFrenquencyList) {
		returnSize = len(wordFrenquencyList)
	}

	out := make([]string, returnSize)
	for i, wf := range wordFrenquencyList[0:returnSize] {
		out[i] = wf.word
	}

	return out
}
