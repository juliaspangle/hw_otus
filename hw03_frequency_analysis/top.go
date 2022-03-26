package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordOccurrences struct {
	word  string
	count int
}

func Top10(inputString string) []string {
	// Place your code here.
	if len(inputString) == 0 {
		return []string{}
	}
	var words []string
	wordsMapOccurrences := map[string]int{}
	wc := make([]wordOccurrences, 0)
	outputStrings := make([]string, 0)

	words = strings.Fields(inputString)
	for _, word := range words {
		wordsMapOccurrences[word]++
	}

	for word, count := range wordsMapOccurrences {
		wc = append(wc, wordOccurrences{word, count})
	}

	sort.Slice(wc, func(i, j int) bool {
		if wc[i].count > wc[j].count {
			return true
		}
		if wc[i].count < wc[j].count {
			return false
		}
		return wc[i].word < wc[j].word
	})

	for _, elem := range wc {
		outputStrings = append(outputStrings, elem.word)
	}

	if topLength := 10; len(outputStrings) > topLength {
		outputStrings = outputStrings[:topLength]
	}

	return outputStrings
}
