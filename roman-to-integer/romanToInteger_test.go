package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	parameters string
	expected   int
}

func TestRomanToInt(t *testing.T) {
	var cases = []Case{
		// Simple case
		Case{
			parameters: "II",
			expected:   2,
		},
		Case{
			parameters: "XII",
			expected:   12,
		},
		Case{
			parameters: "XXVII",
			expected:   27,
		},
		// Need substruction
		Case{
			parameters: "IV",
			expected:   4,
		},
		Case{
			parameters: "IX",
			expected:   9,
		},
		Case{
			parameters: "XL",
			expected:   40,
		},
		Case{
			parameters: "XC",
			expected:   90,
		},
		Case{
			parameters: "CD",
			expected:   400,
		},
		Case{
			parameters: "CM",
			expected:   900,
		},
		// Complex case
		Case{
			parameters: "LVIII",
			expected:   58,
		},
		Case{
			parameters: "MCMXCIV",
			expected:   1994,
		},
		Case{
			parameters: "III",
			expected:   3,
		},
	}
	for _, c := range cases {
		result := romanToInt(c.parameters)
		assert.Equal(t, c.expected, result)
	}
}
