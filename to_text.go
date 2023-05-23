package num2text

import (
	"regexp"
	"strings"
)

var (
	wordsSingular = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tens          = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tensPlace     = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

func ConvertToText(num int) string {
	var result string
	if num == 0 {
		return wordsSingular[0]
	}

	if num < 0 {
		result += "negative "
		num = -num
	}

	// Process billions place
	billions := num / 1000000000
	if billions > 0 {
		result += convertThreeDigitNumber(billions) + " billion "
		num %= 1000000000
	}

	// Process millions place
	millions := num / 1000000
	if millions > 0 {
		result += convertThreeDigitNumber(millions) + " million "
		num %= 1000000
	}

	// Process thousands place
	thousands := num / 1000
	if thousands > 0 {
		result += convertThreeDigitNumber(thousands) + " thousand "
		num %= 1000
	}

	// Process remaining hundreds and tens place
	if num > 0 {
		result += convertThreeDigitNumber(num)
	}
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(strings.TrimSpace(result), " ")
}

func convertThreeDigitNumber(num int) string {
	result := ""

	// Process hundreds place
	hundreds := num / 100
	if hundreds > 0 {
		result += wordsSingular[hundreds] + " hundred "
		num %= 100
	}

	// Process tens place
	if num >= 10 && num <= 19 {
		result += tens[num-10] + " "
		return result
	} else if num >= 20 {
		tens := num / 10
		result += tensPlace[tens] + " "
		num %= 10
	}

	// Process ones place
	if num > 0 {
		result += wordsSingular[num] + " "
	}

	return result
}
