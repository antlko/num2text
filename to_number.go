package num2text

import (
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type tagOptions struct {
	tensMultipl int
}

var (
	tagNames = map[string]tagOptions{
		"hundred":  {tensMultipl: 100},
		"thousand": {tensMultipl: 1000},
		"million":  {tensMultipl: 10000},
		"billion":  {tensMultipl: 100000},
	}
	ordinalSuffTagsSet = map[string]bool{
		"th": true,
		"st": true,
		"nd": true,
		"rd": true,
	}
	ordinalSingl = map[string]int{
		"first":   1,
		"second":  2,
		"third":   3,
		"fifth":   5,
		"eighth":  8,
		"twelfth": 12,
	}
)

func ConvertToNumber(numWord string) int {
	numWord = strings.ToLower(numWord)
	if containsNumbers(numWord) {
		return parseOrdinalToNumber(numWord)
	}

	if len(numWord) > 2 && ordinalSuffTagsSet[numWord[len(numWord)-2:]] {
		return parseFullOrdinalNumber(numWord)
	}

	arr := append(wordsSingular, append(tens, tensPlace...)...)
	if slices.ContainsFunc(arr, func(wordSing string) bool {
		return strings.Contains(numWord, wordSing)
	}) {
		return parseTextNumber(numWord)
	}
	return 0
}

func parseFullOrdinalNumber(numWord string) int {
	words := strings.Split(numWord, " ")

	var out int
	for _, word := range words {
		// ordinal Exept
		if val, ok := ordinalSingl[word]; ok {
			out += val
			continue
		}

		// ordinal with suff
		cutEndWord := word[len(word)-2:]
		cutStartWord := word[:len(word)-2]
		if ordinalSuffTagsSet[cutEndWord] && tagNames[word].tensMultipl == 0 {
			out += checkAndGetNumberResult(cutStartWord)
			continue
		}

		// tags
		if t, ok := tagNames[word]; ok {
			out *= t.tensMultipl
			continue
		}

		out += checkAndGetNumberResult(word)
	}
	return out
}

func checkAndGetNumberResult(word string) int {
	// single
	if idx := slices.IndexFunc(wordsSingular, func(w string) bool {
		return word == w
	}); idx != -1 {
		return idx
	}

	// tens
	if idx := slices.IndexFunc(tensPlace, func(w string) bool {
		return word == w
	}); idx != -1 {
		return idx * 10
	}

	// tens in ordinal
	if idx := slices.IndexFunc(tensWords, func(w string) bool {
		return word == w
	}); idx != -1 {
		return idx * 10
	}

	// 10 >= word < 20
	if idx := slices.IndexFunc(tens, func(w string) bool {
		return word == w
	}); idx != -1 {
		return idx + 10
	}
	return 0
}

func parseOrdinalToNumber(ordinal string) int {
	if len(ordinal) > 2 {
		parseInt, _ := strconv.ParseInt(ordinal[:len(ordinal)-2], 10, 64)
		return int(parseInt)
	}
	return 0
}

func containsNumbers(numWord string) bool {
	if len(numWord) > 0 {
		return rune(numWord[0])-'a' < 0
	}
	return false
}

func parseTextNumber(numWord string) int {
	words := strings.Split(numWord, " ")

	var out int
	for _, word := range words {
		// single
		if idx := slices.IndexFunc(wordsSingular, func(w string) bool {
			return word == w
		}); idx != -1 {
			out += idx
			continue
		}

		// tens
		if idx := slices.IndexFunc(tensPlace, func(w string) bool {
			return word == w
		}); idx != -1 {
			out += idx * 10
			continue
		}

		// 10 >= word < 20
		if idx := slices.IndexFunc(tens, func(w string) bool {
			return word == w
		}); idx != -1 {
			out += idx + 10
			continue
		}

		// tags
		if t, ok := tagNames[word]; ok {
			out *= t.tensMultipl
			continue
		}
	}
	return out
}
