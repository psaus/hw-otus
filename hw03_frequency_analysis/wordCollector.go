package hw03frequencyanalysis

type WordCollector struct {
	list map[string]*wordFrenquency
}

func (w *WordCollector) GetWordFrenquencySorter() *wordFrenquencySorter {
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

func NewWordCollector() *WordCollector {
	return &WordCollector{
		list: make(map[string]*wordFrenquency),
	}
}

func (w *WordCollector) WordUp(value string) {
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
