package num2text

import "strings"

var (
	suffixes = map[string]string{
		"one":    "first",
		"two":    "second",
		"three":  "third",
		"five":   "fifth",
		"eight":  "eighth",
		"twelve": "twelfth",
	}
	tensWords = []string{"", "", "twentie", "thirtie", "fortie", "fiftie", "sixtie", "seventie", "eightie", "ninetie"}
)

// A function that converts a digit to its word form
func digitToWord(digit int) string {
	return wordsSingular[digit]
}

// A function that converts a number from 10 to 19 to its word form
func teenToWord(number int) string {
	return tens[number-10]
}

// A function that converts a number from 20 to 99 to its word form
func tensToWord(number int) string {
	tens := number / 10
	ones := number % 10
	if ones == 0 {
		return tensWords[tens]
	} else {
		return tensPlace[tens] + " " + digitToWord(ones)
	}
}

// A function that converts a number from 100 to 999 to its word form
func hundredsToWord(number int) string {
	hundreds := number / 100
	rest := number % 100
	if rest == 0 {
		return digitToWord(hundreds) + " hundred"
	} else {
		return digitToWord(hundreds) + " hundred " + numberToWord(rest)
	}
}

// A function that converts a number from 1000 to 999999 to its word form
func thousandsToWord(number int) string {
	thousands := number / 1000
	rest := number % 1000
	if rest == 0 {
		return numberToWord(thousands) + " thousand"
	} else {
		return numberToWord(thousands) + " thousand " + numberToWord(rest)
	}
}

// A function that converts any number from 0 to 999999 to its word form
func numberToWord(number int) string {
	var negative string
	if number < 0 {
		negative = "minus "
		number *= -1
	}
	if number < 10 {
		return negative + digitToWord(number)
	} else if number < 20 {
		return negative + teenToWord(number)
	} else if number < 100 {
		return negative + tensToWord(number)
	} else if number < 1000 {
		return negative + hundredsToWord(number)
	} else {
		return negative + thousandsToWord(number)
	}
}

// A function that adds the ordinal suffix to a word
func addOrdinalSuffix(word string) string {
	s := getNumSuffix(word)
	value, ok := suffixes[s]
	if ok {
		lastIdx := strings.LastIndex(word, s)
		return word[:lastIdx] + value
	}
	return word + "th"
}

func getNumSuffix(word string) string {
	for r := len(word) - 1; r >= 0; r-- {
		if word[r] == ' ' {
			return strings.TrimSpace(word[r:])
		}
	}
	return word
}

// ConvertToOrdinalFull A function that converts a number to an ordinal in full form
func ConvertToOrdinalFull(number int) string {
	word := numberToWord(number)
	ordinal := addOrdinalSuffix(word)
	return ordinal
}
