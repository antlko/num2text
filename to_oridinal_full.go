package num2text

import "strings"

var (
	suffixes = map[string]string{
		"one":     "first",
		"two":     "second",
		"three":   "third",
		"five":    "fifth",
		"eight":   "eighth",
		"twelve":  "twelfth",
		"-one":    "-first",
		"-two":    "-second",
		"-three":  "-third",
		"-five":   "-fifth",
		"-eight":  "-eighth",
		"-twelve": "-twelfth",
	}
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
	words := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	tens := number / 10
	ones := number % 10
	if ones == 0 {
		return words[tens-2]
	} else {
		return words[tens-2] + "-" + digitToWord(ones)
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

// A function that converts a number from 1000000 to 999999999 to its word form
func millionsToWord(number int) string {
	millions := number / 1000000
	rest := number % 1000000
	if rest == 0 {
		return numberToWord(millions) + " million"
	} else {
		return numberToWord(millions) + " million " + numberToWord(rest)
	}
}

// A function that converts any number from 0 to 999999999 to its word form
func numberToWord(number int) string {
	var negative = ""
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
	} else if number < 1000000 {
		return negative + thousandsToWord(number)
	} else {
		return negative + millionsToWord(number)
	}
}

// A function that adds the ordinal suffix to a word
func addOrdinalSuffix(word string) string {
	s, _ := getNumSuffix(word)
	value, ok := suffixes[s]
	if ok {
		return strings.Replace(word, s, value, -1)
	}
	// If the word contains a dash and the last part ends with one of the keys in the suffixes map, replace it with the corresponding value
	if strings.Contains(word, "-") && strings.HasSuffix(strings.Split(word, "-")[1], word) {
		return word + value
	}
	return word + "th"
}

func getNumSuffix(word string) (string, int) {
	for r := len(word) - 1; r >= 0; r-- {
		if word[r] == ' ' || word[r] == '-' {
			return strings.TrimSpace(word[r:]), r
		}
	}
	return word, 0
}

// ConvertToOrdinalFull A function that converts a number to an ordinal in full form
func ConvertToOrdinalFull(number int) string {
	word := numberToWord(number)
	ordinal := addOrdinalSuffix(word)
	return ordinal
}
