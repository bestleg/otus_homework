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
	sort.Slice(tempStr2, func(i, j int) bool {
		return tempStr2[i] < tempStr2[j]
	})
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
	fmt.Println(m["Кристофер"])
	return nil
}
