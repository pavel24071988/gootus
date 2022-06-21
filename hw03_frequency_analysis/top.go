package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string, taskWithAsteriskIsCompleted bool) []string {
	var resultSlice []string
	if taskWithAsteriskIsCompleted {
		re := regexp.MustCompile(`[^0-9a-zA-Zа-яА-Я-\s]`)
		text = strings.ToLower(re.ReplaceAllString(text, ""))

		prepareWordsByIds, sliceOfCounts := prepareData(text)

		for _, countOfWords := range sliceOfCounts {
			for _, word := range prepareWordsByIds[countOfWords] {
				if len(resultSlice) < 10 && word != `-` {
					resultSlice = append(resultSlice, word)
				}
			}
		}

		return resultSlice
	}

	prepareWordsByIds, sliceOfCounts := prepareData(text)

	for _, countOfWords := range sliceOfCounts {
		for _, word := range prepareWordsByIds[countOfWords] {
			if len(resultSlice) < 10 {
				resultSlice = append(resultSlice, word)
			}
		}
	}

	return resultSlice
}

func prepareData(text string) (map[int][]string, []int) {
	sliceOfCounts := make([]int, 0)
	prepareWordsByIds := make(map[int][]string)
	prepareWords := make(map[string]int)

	words := strings.Fields(text)
	for _, word := range words {
		if _, exist := prepareWords[word]; !exist {
			prepareWords[word] = 0
		}
		prepareWords[word]++
	}

	for prepareWord, count := range prepareWords {
		prepareWordsByIds[count] = append(prepareWordsByIds[count], prepareWord)
	}

	for count, wordsSlice := range prepareWordsByIds {
		sortStringSlice(wordsSlice)
		sliceOfCounts = append(sliceOfCounts, count)
	}

	sort.Slice(sliceOfCounts, func(i, j int) bool {
		return sliceOfCounts[i] > sliceOfCounts[j]
	})

	return prepareWordsByIds, sliceOfCounts
}

func sortStringSlice(wordsSlice []string) []string {
	sort.Slice(wordsSlice, func(i, j int) bool {
		return wordsSlice[i] < wordsSlice[j]
	})
	return wordsSlice
}
