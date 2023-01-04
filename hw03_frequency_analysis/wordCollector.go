package hw03frequencyanalysis

type wordCollector struct {
	list map[string]*wordFrenquency
}

func (w *wordCollector) GetWordFrenquencySorter() *wordFrenquencySorter {
	wordListSize := len(w.list)

	if wordListSize == 0 {
		return &wordFrenquencySorter{
			wordStats: []wordFrenquency{},
		}
	}

	list := make([]wordFrenquency, 0, wordListSize)
	for _, value := range w.list {
		list = append(list, *value)
	}

	return &wordFrenquencySorter{
		wordStats: list,
	}
}

func NewWordCollector() *wordCollector {
	return &wordCollector{
		list: make(map[string]*wordFrenquency),
	}
}

func (w *wordCollector) wordUp(value string) {
	wf, ok := w.list[value]

	if !ok {
		w.list[value] = &wordFrenquency{
			counter: 1,
			word:    value,
		}
	} else {
		wf.counter++
	}
}
