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

type ByFrenquency []*wordFrenquency

func (f ByFrenquency) Len() int      { return len(f) }
func (f ByFrenquency) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f ByFrenquency) Less(i, j int) bool {
	if f[i].counter == f[j].counter {
		return f[i].word < f[j].word
	}

	return f[i].counter > f[j].counter
}

var cleanUpWordRegexp = regexp.MustCompile(`[^a-zA-Z\\u0430-яА-Я-]+`)

func Top10(text string) []string {
	collector := make(map[string]int)

	for _, value := range strings.Fields(text) {
		value = strings.Trim(value, "-")
		if len(value) == 0 {
			continue
		}

		value := strings.ToLower(cleanUpWordRegexp.ReplaceAllString(value, ""))
		if _, ok := collector[value]; !ok {
			collector[value] = 1
		} else {
			collector[value]++
		}
	}

	i := 0
	listSize := len(collector)
	wordFrenquencyList := make([]*wordFrenquency, listSize)
	for w, c := range collector {
		wordFrenquencyList[i] = &wordFrenquency{c, w}
		i++
	}

	sort.Sort(ByFrenquency(wordFrenquencyList))
	returnSize := 10
	if listSize < 10 {
		returnSize = listSize
	}

	out := make([]string, returnSize)
	for i, wf := range wordFrenquencyList[0:returnSize] {
		out[i] = wf.word
	}
	return out
}
