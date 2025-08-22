package entity

import (
	"slices"
)

type (
	Vocabulary struct {
		words map[string]int
	}

	Word struct {
		Text  string
		Count int
	}
)

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		words: make(map[string]int),
	}
}

func (w *Vocabulary) AddOrInc(word string) {
	w.words[word]++
}

func (w *Vocabulary) GetUniqueCount() (c int) {
	for _, v := range w.words {
		if v == 1 {
			c++
		}
	}

	return c
}

func (w *Vocabulary) GetMostFrequent(num int) []Word {
	pairs := make([]Word, 0, len(w.words))

	for k, v := range w.words {
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

	return pairs[:num]
}
