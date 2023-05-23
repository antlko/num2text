package num2text

import "strconv"

var (
	ordinalSuffixes = []string{"th", "st", "nd", "rd", "th", "th", "th", "th", "th", "th"}
)

func ConvertToOrdinal(num int) string {
	return strconv.Itoa(num) + ordinalSuffix(num)
}

func ordinalSuffix(num int) string {
	if num >= 10 && num <= 20 {
		return "th"
	}

	lastDigit := num % 10
	if lastDigit >= 1 && lastDigit <= 3 {
		if num%100/10 != 1 {
			switch lastDigit {
			case 1:
				return "st"
			case 2:
				return "nd"
			case 3:
				return "rd"
			default:
				return ordinalSuffixes[lastDigit]
			}
		}
	}
	return "th"
}
