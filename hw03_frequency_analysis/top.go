package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFrequency struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	tempStr := strings.Fields(text)
	m := make(map[string]int)
	for _, word := range tempStr {
		m[word]++
	}

	resultStruct := sortValues(m)
	result := make([]string, 0, 10)
	for i, v := range resultStruct {
		result = append(result, v.Word)
		if i == 9 {
			break
		}
	}
	return result
}

func sortValues(m map[string]int) []wordFrequency {
	valuesStruct := make([]wordFrequency, 0, len(m))
	for word, count := range m {
		valuesStruct = append(valuesStruct, wordFrequency{
			Word:  word,
			Count: count,
		})
	}
	sort.SliceStable(valuesStruct, func(i, j int) bool {
		return valuesStruct[i].Word < valuesStruct[j].Word
	})
	sort.SliceStable(valuesStruct, func(i, j int) bool {
		return valuesStruct[i].Count > valuesStruct[j].Count
	})
	return valuesStruct
}
