package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	tempStr := strings.Fields(text)
	m := make(map[string]int)

	for _, word := range tempStr {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] = m[word] + 1
		}
	}

	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	RemoveDuplicates(&values)
	result := make([]string, 0, len(m))
	for _, chislo := range values {
		var tempRes []string
		for k, v := range m {
			if v == chislo {
				tempRes = append(tempRes, k)
			}
		}
		sort.Strings(tempRes)
		result = append(result, tempRes...)
	}
	return result[0:10]
}

func RemoveDuplicates(xs *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
