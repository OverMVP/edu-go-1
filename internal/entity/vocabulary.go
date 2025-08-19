package entity

import "slices"

type (
	Vocabulary map[string]int

	Word struct {
		Text  string
		Count int
	}
)

func (w Vocabulary) AddOrInc(word string) {
	if _, ok := w[word]; !ok {
		w[word] = 1
	} else {
		w[word]++
	}
}

func (w Vocabulary) GetUniqueCount() (count int64) {
	for _, v := range w {
		if v == 1 {
			count++
		}
	}

	return
}

func (w Vocabulary) GetMostFrequent(num int) []Word {
	pairs := make([]Word, 0, len(w))

	for k, v := range w {
		pairs = append(pairs, Word{Text: k, Count: v})
	}

	slices.SortStableFunc(pairs, func(a, b Word) int {
		switch {

		case a.Count > b.Count:
			return -1

		case a.Count < b.Count:
			return 1

		default:
			return 0
		}
	})

	if len(pairs) <= num {
		return pairs
	}

	return pairs[0:num]
}
