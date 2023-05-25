package num2text_test

import (
	"github.com/antlko/num2text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToText(t *testing.T) {
	tests := []struct {
		num    int
		result string
	}{
		{
			1,
			"one",
		},
		{
			0,
			"zero",
		},
		{
			2,
			"two",
		},
		{
			3,
			"three",
		},
		{
			100,
			"one hundred",
		},
		{
			54,
			"fifty four",
		},
		{
			13,
			"thirteen",
		},
		{
			66,
			"sixty six",
		},
		{
			999,
			"nine hundred ninety nine",
		},
		{
			999999,
			"nine hundred ninety nine thousand nine hundred ninety nine",
		},
		{
			-112,
			"negative one hundred twelve",
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, num2text.ConvertToText(test.num))
		})
	}
}

func TestNumberToOrdinal(t *testing.T) {
	tests := []struct {
		num    int
		result string
	}{
		{
			1,
			"1st",
		},
		{
			2,
			"2nd",
		},
		{
			3,
			"3rd",
		},
		{
			100,
			"100th",
		},
		{
			54,
			"54th",
		},
		{
			13,
			"13th",
		},
		{
			66,
			"66th",
		},
		{
			999,
			"999th",
		},
		{
			999999,
			"999999th",
		},
		{
			-112,
			"-112th",
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, num2text.ConvertToOrdinal(test.num))
		})
	}
}

func TestNumberToOrdinalShortForm(t *testing.T) {
	tests := []struct {
		num    int
		result string
	}{
		{
			1,
			"1st",
		},
		{
			2,
			"2nd",
		},
		{
			3,
			"3rd",
		},
		{
			100,
			"100th",
		},
		{
			54,
			"54th",
		},
		{
			13,
			"13th",
		},
		{
			66,
			"66th",
		},
		{
			999,
			"999th",
		},
		{
			999999,
			"999999th",
		},
		{
			-112,
			"-112th",
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, num2text.ConvertToOrdinal(test.num))
		})
	}
}

func TestNumberToOrdinalFullForm(t *testing.T) {
	tests := []struct {
		num    int
		result string
	}{
		{
			0,
			"zeroth",
		},
		{
			1,
			"first",
		},
		{
			2,
			"second",
		},
		{
			3,
			"third",
		},
		{
			100,
			"one hundredth",
		},
		{
			54,
			"fifty fourth",
		},
		{
			13,
			"thirteenth",
		},
		{
			66,
			"sixty sixth",
		},
		{
			999,
			"nine hundred ninety nineth",
		},
		{
			999999,
			"nine hundred ninety nine thousand nine hundred ninety nineth",
		},
		{
			-112,
			"minus one hundred twelfth",
		},
		{
			55,
			"fifty fifth",
		},
		{
			105,
			"one hundred fifth",
		},
		{
			101,
			"one hundred first",
		},
		{
			20,
			"twentieth",
		},
		{
			30,
			"thirtieth",
		},
		{
			50,
			"fiftieth",
		},
		{
			80,
			"eightieth",
		},
		{
			180,
			"one hundred eightieth",
		},
		{
			10,
			"tenth",
		},
		{
			1055,
			"one thousand fifty fifth",
		},
		{
			1000,
			"one thousandth",
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, num2text.ConvertToOrdinalFull(test.num))
		})
	}
}

func TestTexToNumber(t *testing.T) {
	tests := []struct {
		num   int
		title string
	}{
		{
			0,
			"zero",
		},
		{
			1,
			"one",
		},
		{
			9,
			"nine",
		},
		{
			11,
			"eleven",
		},
		{
			12,
			"twelve",
		},
		{
			15,
			"fifteen",
		},
		{
			21,
			"twenty one",
		},
		{
			44,
			"44",
		},
		{
			79,
			"79",
		},
		{
			99,
			"99",
		},
		{
			101,
			"101",
		},
		{
			123,
			"123",
		},
		{
			10,
			"10",
		},
		{
			19,
			"19",
		},
		{
			20,
			"20",
		},
		{
			30,
			"30",
		},
		{
			70,
			"70",
		},
		{
			87,
			"87",
		},
		{
			112,
			"112",
		},
		{
			1012,
			"1012",
		},
		{
			10012,
			"10012",
		},
		{
			999099,
			"999099",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			text := num2text.ConvertToText(test.num)
			assert.Equal(t, test.num, num2text.ConvertToNumber(text))
		})
	}
}

func TestOrdinalToNumber(t *testing.T) {
	tests := []struct {
		num   int
		title string
	}{
		{
			0,
			"zero",
		},
		{
			1,
			"one",
		},
		{
			9,
			"nine",
		},
		{
			11,
			"eleven",
		},
		{
			12,
			"twelve",
		},
		{
			15,
			"fifteen",
		},
		{
			21,
			"twenty one",
		},
		{
			44,
			"44",
		},
		{
			79,
			"79",
		},
		{
			99,
			"99",
		},
		{
			101,
			"101",
		},
		{
			123,
			"123",
		},
		{
			10,
			"10",
		},
		{
			19,
			"19",
		},
		{
			20,
			"20",
		},
		{
			30,
			"30",
		},
		{
			70,
			"70",
		},
		{
			87,
			"87",
		},
		{
			112,
			"112",
		},
		{
			1012,
			"1012",
		},
		{
			10012,
			"10012",
		},
		{
			999099,
			"999099",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			text := num2text.ConvertToOrdinal(test.num)
			assert.Equal(t, test.num, num2text.ConvertToNumber(text))
		})
	}
}

func TestFullOrdinalToNumber(t *testing.T) {
	tests := []struct {
		num   int
		title string
	}{
		{
			8,
			"8",
		},
		{
			88,
			"88",
		},
		{
			0,
			"zero",
		},
		{
			1,
			"one",
		},
		{
			9,
			"nine",
		},
		{
			11,
			"eleven",
		},
		{
			12,
			"twelve",
		},
		{
			15,
			"fifteen",
		},
		{
			21,
			"twenty one",
		},
		{
			44,
			"44",
		},
		{
			79,
			"79",
		},
		{
			99,
			"99",
		},
		{
			101,
			"101",
		},
		{
			123,
			"123",
		},
		{
			10,
			"10",
		},
		{
			19,
			"19",
		},
		{
			20,
			"20",
		},
		{
			30,
			"30",
		},
		{
			70,
			"70",
		},
		{
			87,
			"87",
		},
		{
			112,
			"112",
		},
		{
			1012,
			"1012",
		},
		{
			10012,
			"10012",
		},
		{
			999099,
			"999099",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			text := num2text.ConvertToOrdinalFull(test.num)
			assert.Equal(t, test.num, num2text.ConvertToNumber(text))
		})
	}
}
