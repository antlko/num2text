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
			"fifty-fourth",
		},
		{
			13,
			"thirteenth",
		},
		{
			66,
			"sixty-sixth",
		},
		{
			999,
			"nine hundred ninety-nineth",
		},
		{
			999999,
			"nine hundred ninety-nine thousand nine hundred ninety-nineth",
		},
		{
			-112,
			"minus one hundred twelfth",
		},
		{
			55,
			"fifty-fifth",
		},
		{
			105,
			"one hundred fifth",
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, num2text.ConvertToOrdinalFull(test.num))
		})
	}
}
