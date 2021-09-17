package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(text string) []string {
	//tempStr := strings.Replace(text,"\t"," ",-1)
	//strings.Fields(tempStr," ")
	//tempStr2 := strings.Split(tempStr," ")
	tempStr2 := strings.Fields(text)
	//sort.Strings(tempStr2)
	//sort.Slice(tempStr2, func(i, j int) bool {
	//	return tempStr2[i] < tempStr2[j]
	//})
	//for _, word := range tempStr2 {
	//re := regexp.MustCompile(`[[:blank:]]|[[:space:]]`)
	//word = re.ReplaceAllString(word,"")
	//fmt.Println(word)
	//}
	m := make(map[string]int)
	for _, word := range tempStr2 {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] = m[word] + 1
		}
	}
	keys := make([]int, 0, len(m))
	for _, v := range m {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	RemoveDuplicates(&keys)
	for _, chislo := range keys {
		for k, v := range m {
			if v == chislo {
				fmt.Println(k, m[k])
			}
		}
	}

	return nil
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
