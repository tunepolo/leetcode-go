package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	parameters int
	expected   bool
}

var testCase = []Case{
	Case{
		parameters: 121,
		expected:   true,
	},
	Case{
		parameters: -121,
		expected:   false,
	},
	Case{
		parameters: 10,
		expected:   false,
	},
	Case{
		parameters: 1221,
		expected:   true,
	},
	Case{
		parameters: 12321,
		expected:   true,
	},
}

func TestIsPalindrome(t *testing.T) {
	cases := testCase

	for _, c := range cases {
		result := isPalindrome(c.parameters)
		assert.Equal(t, c.expected, result)
	}
}

func TestIsPalindrome2(t *testing.T) {
	cases := testCase

	for _, c := range cases {
		result := isPalindrome2(c.parameters)
		assert.Equal(t, c.expected, result)
	}
}
