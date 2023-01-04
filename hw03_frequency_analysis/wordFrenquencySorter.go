package hw03frequencyanalysis

import "sort"

type wordFrenquencySorter struct {
	wordStats []wordFrenquency
}

func (ws *wordFrenquencySorter) GetSortedWords(maxFirstWords int) []string {
	hasMaxWords := maxFirstWords > 0
	wordStatsLen := len(ws.wordStats)

	if !hasMaxWords {
		maxFirstWords = wordStatsLen
	}

	if maxFirstWords == 0 || wordStatsLen == 0 {
		return []string{}
	}

	sort.Sort(ws)
	result := make([]string, 0, maxFirstWords)

	for _, v := range ws.wordStats[0:maxFirstWords] {
		result = append(result, v.word)
	}

	return result
}

func (ws *wordFrenquencySorter) Len() int {
	return len(ws.wordStats)
}

func (ws *wordFrenquencySorter) Swap(i, j int) {
	ws.wordStats[i], ws.wordStats[j] = ws.wordStats[j], ws.wordStats[i]
}

func (ws *wordFrenquencySorter) Less(i, j int) bool {
	if ws.wordStats[i].counter == ws.wordStats[j].counter {
		return ws.wordStats[i].word < ws.wordStats[j].word
	}

	return ws.wordStats[i].counter > ws.wordStats[j].counter
}
