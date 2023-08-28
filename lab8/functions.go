package main 

import (
	"errors"
	"sort"
	"strings"
)

func split(text string, s ...string) ([]string, error) {
	sep := " "
	if len(s) == 1 {
		sep = s[0]
	} else if len(s) > 1 {
		return nil, errors.New("Too many arguments")
	}

	splited := strings.Split(text, sep)
	return splited, nil
}

func splitAndSort(text string, s ...string) ([]string, error) {
	sep := " "
	if len(s) == 1 {
		sep = s[0]
	} else if len(s) > 1 {
		return nil, errors.New("Too many arguments")
	}

	splited := strings.Split(text, sep)
	sort.Strings(splited)

	return splited, nil
}

func getNumberOfNonEmptyLines(text string) int {
	lines := strings.Split(text, "\n")
	numOfNonEmptyLines := 0
	for _, line := range lines {
		if line != "" {
			numOfNonEmptyLines++
		}
	}
	return numOfNonEmptyLines
}

func getNumOfWords(text string) int {
	return len(strings.Split(text, " "))
}

func getNumOfNonWhiteChars(text string) int {
	return len(text) - strings.Count(text, " ") - strings.Count(text, "\t") -
		strings.Count(text, "\n")
}

func getFrequencyForWords(text string) map[string] int {
	mapWordsFrequency := make(map[string] int)
	words := strings.Split(text, " ")
	for _, word := range words {
		mapWordsFrequency[word]++
	}
	return mapWordsFrequency
}

func getSortedPalindroms(text string) []string {
	words := strings.Split(text, " ")
	palindroms := make([]string, 0)
	for _, word := range words {
		palindrom := true
		for i := 0; i < len(word) / 2; i++ {
			if word[i] != word[len(word) - 1 - i] {
				palindrom = false
				break
			}
		}
		if palindrom {
			palindroms = append(palindroms, word)
		}
	}
	sort.Strings(palindroms)
	return palindroms
}